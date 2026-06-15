package main

import (
	"context"
	"log"
	"net/http"
"time"
	"github.com/google/uuid"

	"github.com/cloud-cost-iq/config"
	"github.com/cloud-cost-iq/internals/account"
	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/analytics"
	"github.com/cloud-cost-iq/internals/api"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/db"
	"github.com/cloud-cost-iq/internals/ingestion"
	"github.com/cloud-cost-iq/internals/recommendation"
	"github.com/cloud-cost-iq/internals/worker"
)


func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }
	ctx := context.Background()

	dbConn, err := db.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// db.ResetMigrations(cfg.DatabaseURL)
	db.RunMigrations(cfg.DatabaseURL)
	

	// Repos
	billingRepo := billing.NewRepository(dbConn)
	aggregateRepo := aggregation.NewRepository(dbConn)
	ingestionRepo := ingestion.NewRepository(billingRepo)
	analyticsRepo := analytics.NewRepository(dbConn)
	accountRepo := account.NewRepository(dbConn)
	recommendationRepo := recommendation.NewRepository(dbConn)


	//Services
	billingService := billing.NewService(billingRepo)
	aggregateService := aggregation.NewService(aggregateRepo)
	ingestionService := ingestion.NewService(ingestionRepo)
	analyticsService := analytics.NewService(analyticsRepo)
	accountService := account.NewService(accountRepo)
	recommendationEngine := recommendation.NewEngine(analyticsService)
	recommendationService := recommendation.NewService(recommendationEngine, recommendationRepo)
	// Handlers
	
	billingHandler := billing.NewHandler(billingService)
	aggregationHandler := aggregation.NewHandler(aggregateService)
	analyticsHandler := analytics.NewHandler(analyticsService)
	accountHandler := account.NewHandler(accountService)
	recommendationHandler := recommendation.NewHandler(recommendationService)



	queue := worker.NewQueue(1000)

	processor := worker.NewProcessor(ingestionService)

	pool := worker.NewPool(queue, 5, processor)

	pool.Start(ctx)

	queue.Push(worker.Job{
	InternalID: uuid.New(),

	Payload: ingestion.RawCostEvent{
		Provider: "AWS",

		AwsAccountID: "482836517294",

		Service: "EC2",

		Region: "us-east-1",

		UsageAmount: 10,

		CostAmount: 2.5,

		Currency: "USD",
		UsageStart:  time.Now().AddDate(0, 0, -1), // ← yesterday
        UsageEnd:    time.Now(), 
	},
	
	})


	router := api.NewRouter(billingHandler, aggregationHandler, analyticsHandler, accountHandler, recommendationHandler)



    log.Printf("API listening on :%s", cfg.Port)

    if err = http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatal(err)
    }
}
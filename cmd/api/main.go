package main

import (
	"context"
	"log"
	"net/http"

	// "github.com/google/uuid"

	"github.com/cloud-cost-iq/config"
	"github.com/cloud-cost-iq/internals/account"
	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/analytics"
	"github.com/cloud-cost-iq/internals/api"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/db"
	"github.com/cloud-cost-iq/internals/ingestion"
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
	
	// BillingRepo := billing.NewRepository(dbConn)
	// aggreagteRepo := aggregation.NewRepository(dbConn)
	// ingestionRepo := ingestion.NewRepository(BillingRepo)
	// billingService := billing.NewService(BillingRepo)
	// analyticsRepo := analytics.NewRepository(dbConn)

	// Repos
	billingRepo := billing.NewRepository(dbConn)
	aggregateRepo := aggregation.NewRepository(dbConn)
	ingestionRepo := ingestion.NewRepository(billingRepo)
	analyticsRepo := analytics.NewRepository(dbConn)
	accountRepo := account.NewRepository(dbConn)

	// ingestionService := ingestion.NewService(ingestionRepo)
	// aggregateService := aggregation.NewService(aggreagteRepo)
	// analyticsService := analytics.NewService(analyticsRepo)

	//Services
	billingService := billing.NewService(billingRepo)
	aggregateService := aggregation.NewService(aggregateRepo)
	ingestionService := ingestion.NewService(ingestionRepo)
	analyticsService := analytics.NewService(analyticsRepo)
	accountService := account.NewService(accountRepo)

	// Handlers
	billingHandler := billing.NewHandler(billingService)
	aggregationHandler := aggregation.NewHandler(aggregateService)
	analyticsHandler := analytics.NewHandler(analyticsService)
	accountHandler := account.NewHandler(accountService)



	queue := worker.NewQueue(1000)

	processor := worker.NewProcessor(ingestionService)

	pool := worker.NewPool(queue, 5, processor)

	pool.Start(ctx)

	// queue.Push(worker.Job{
	// InternalID: uuid.New(),

	// Payload: ingestion.RawCostRecord{
	// 	Provider: "AWS",

	// 	AwsAccountID: uuid.NewString(),

	// 	Service: "EC2",

	// 	Region: "us-east-1",

	// 	UsageAmount: 10,

	// 	CostAmount: 2.5,

	// 	Currency: "USD",
	// },
	// })

    // router := api.NewRouter(billingService, aggregateService, analyticsService)

	router := api.NewRouter(billingHandler, aggregationHandler, analyticsHandler, accountHandler)



    log.Printf("API listening on :%s", cfg.Port)

    if err = http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatal(err)
    }
}
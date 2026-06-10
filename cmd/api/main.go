package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cloud-cost-iq/config"
	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/api"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/db"
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

	db.RunMigrations(cfg.DatabaseURL)
	BillingRepo := billing.NewRepository(dbConn)
	aggreagteRepo := aggregation.NewRepository(dbConn)
	billingService := billing.NewService(BillingRepo)
	aggreagteService := aggregation.NewService(aggreagteRepo)

    router := api.NewRouter(billingService, aggreagteService)

    log.Printf("API listening on :%s", cfg.Port)

    if err = http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatal(err)
    }
}
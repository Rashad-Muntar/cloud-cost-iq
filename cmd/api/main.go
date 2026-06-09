package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cloud-cost-iq/config"
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
	repo := billing.NewRepository(dbConn)
	_ = billing.NewService(repo)

    router := api.NewRouter()

    log.Printf("API listening on :%s", cfg.Port)

    if err = http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatal(err)
    }
}
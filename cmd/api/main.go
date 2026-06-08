package main

import (
    "log"
    "net/http"

    "github.com/cloud-cost-iq/config"
    "github.com/cloud-cost-iq/internals/api"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    router := api.NewRouter()

    log.Printf("API listening on :%s", cfg.Port)

    if err = http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatal(err)
    }
}
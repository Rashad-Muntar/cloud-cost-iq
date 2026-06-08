package main

import (
	"log"
	"net/http"

	"github.com/cloud-cost-iq/internals/api"
)

func main() {
	router := api.NewRouter()

	log.Println("API listening on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
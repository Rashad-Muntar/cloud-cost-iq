package api

import (
	"net/http"

	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(billingService *billing.Service, aggregationService *aggregation.Service) http.Handler {
	r := chi.NewRouter()
	handler := NewHandlers(billingService, aggregationService) // ← pass billing service here
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	r.Post("/costs", handler.CreateCost)
	r.Get("/costs/summary/daily", handler.GetDailySummary)
	r.Post("/aggregation/run", handler.ExcuteDailySummaryJob)
	r.Post("/accounts", handler.CreateAccount)
	r.Get("/accounts", handler.ListAccounts)
	r.Get("/account", handler.GetByAWSAccountID)
	return r
}
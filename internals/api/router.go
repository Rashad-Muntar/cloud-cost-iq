package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/cloud-cost-iq/internals/billing"
)

func NewRouter(billingService *billing.Service) http.Handler {
	r := chi.NewRouter()
	handler := NewHandlers(billingService) // ← pass billing service here
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
	return r
}
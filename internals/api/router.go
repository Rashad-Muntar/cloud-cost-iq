package api

import (
	"net/http"

	"github.com/cloud-cost-iq/internals/account"
	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/analytics"
	"github.com/cloud-cost-iq/internals/anomaly"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/recommendation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(billingHandler *billing.Handler, aggregationHandler *aggregation.Handler, analyticsHandler *analytics.Handler, accountHandler *account.Handler, recommendation *recommendation.Handler, anomaly *anomaly.Handler) http.Handler {
	r := chi.NewRouter()
	handler := NewHandlers(billingHandler, aggregationHandler, analyticsHandler, accountHandler, recommendation, anomaly)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	r.Post("/costs", handler.billingHandler.CreateCost)
	r.Get("/costs/summary/daily", handler.billingHandler.GetDailySummary)
	r.Post("/aggregation/run", handler.aggregationHandler.ExcuteDailySummaryJob)
	r.Post("/accounts", handler.accountHandler.CreateAccount)
	r.Get("/accounts", handler.accountHandler.ListAccounts)
	r.Get("/account", handler.accountHandler.GetByAWSAccountID)
	r.Get("/analytics/summary", handler.analyticsHandler.Summary)
	r.Post("/recommendations/generate", handler.recommendationHandler.Generate,
)
r.Post("/anomaly/run", handler.anomalyHandler.Run,)
	return r
}
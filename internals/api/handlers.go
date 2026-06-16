package api

import (
	"github.com/cloud-cost-iq/internals/account"
	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/analytics"
	"github.com/cloud-cost-iq/internals/anomaly"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/recommendation"
)

type Handlers struct {
	billingHandler *billing.Handler
	aggregationHandler *aggregation.Handler
	accountHandler *account.Handler
	analyticsHandler *analytics.Handler
	recommendationHandler *recommendation.Handler
	anomalyHandler *anomaly.Handler
}

func NewHandlers(billingHandler *billing.Handler, aggregationHandler *aggregation.Handler, analyticsHandler *analytics.Handler, accountHandler *account.Handler, recommendationHandler *recommendation.Handler, anomalyHandler *anomaly.Handler) *Handlers {
	return &Handlers{billingHandler: billingHandler, aggregationHandler: aggregationHandler, analyticsHandler: analyticsHandler, accountHandler:accountHandler, recommendationHandler:recommendationHandler, anomalyHandler: anomalyHandler}
}
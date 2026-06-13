package api

import (

	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/account"
	"github.com/cloud-cost-iq/internals/analytics"
)

type Handlers struct {
	billingHandler *billing.Handler
	aggregationHandler *aggregation.Handler
	accountHandler *account.Handler
	analyticsHandler *analytics.Handler
}

func NewHandlers(billingHandler *billing.Handler, aggregationHandler *aggregation.Handler, analyticsHandler *analytics.Handler, accountHandler *account.Handler) *Handlers {
	return &Handlers{billingHandler: billingHandler, aggregationHandler: aggregationHandler, analyticsHandler: analyticsHandler, accountHandler:accountHandler}
}
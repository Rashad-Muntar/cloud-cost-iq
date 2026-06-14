package recommendation

import (
	"context"

	"github.com/cloud-cost-iq/internals/analytics"
)

type Engine struct {
	analytics *analytics.Service
}

func NewEngine(
	a *analytics.Service,
) *Engine {

	return &Engine{
		analytics: a,
	}
}

func (e *Engine) Generate(
	ctx context.Context,
	accountID string,
)(
	[]Recommendation,
	error,
){

	summary,
	err :=
	e.analytics.
	GetCostSummary(
		ctx,

		analytics.CostQuery{
			AwsAccountID: accountID,
		},
	)

	if err != nil {
		return nil,
		err
	}

	var output []Recommendation
	for _, service := range summary.Breakdown {
		if service.Service == "EC2" && service.Cost < 20 {
			output = append(output, Recommendation{
					AccountID: accountID,
					Type: IdleCompute,
					Title: "Low EC2 utilization",
					Description: "review EC2 usage",
					EstimatedMonthlySavings: service.Cost * 0.3,
					Severity: Medium,
					Status: Open,
				},
			)
		}
	}

	return output,nil
}


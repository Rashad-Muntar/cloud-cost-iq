package recommendation

import (
	"context"
	"fmt"
	"time"

	"github.com/cloud-cost-iq/internals/analytics"
	"github.com/cloud-cost-iq/internals/shared"
)

type Engine struct {
	analytics *analytics.Service
}

func NewEngine(
	analytics *analytics.Service,
) *Engine {

	return &Engine{
		analytics: analytics,
	}
}

func (e *Engine) Generate(
	ctx context.Context,
	accountID string,
	from string,
	to string,
)(
	[]Recommendation,
	error,
){
	
	var parsedFrom time.Time
	var parsedTo time.Time

if from != "" {
    parsed, err := shared.ParseDate(from)
    if err != nil {
      return nil, err
    }
    parsedFrom = parsed  
}

if to != "" {
    parsed, err := shared.ParseDate(to)
      if err != nil {
      return nil, err
    }
    parsedTo = shared.EndOfDay(parsed)
}

	fmt.Println("Reco Service hit", parsedFrom, parsedTo)
	summary, err := e.analytics.GetCostSummary(ctx, analytics.CostQuery{
			AwsAccountID: accountID,
			From: parsedFrom,
			To: parsedTo,
		},
	)
	fmt.Println("Summary", summary)
	if err != nil {
		return nil, err
	}

	var output []Recommendation
	for _, service := range summary.Breakdown {
		fmt.Println("Service in", service)
		if service.Service == "EC2" && service.Cost > 20 {
			output = append(output, Recommendation{
					AccountID: summary.InternalID,
					Type: IdleCompute,
					Title: "Low EC2 utilization",
					Description: "review EC2 usage",
					EstimatedMonthlySavings: service.Cost * 0.3,
					Severity: Medium,
					Status: Open,
				},
			
			)
			fmt.Println(output)
		}
	}
	
	return output,nil
}


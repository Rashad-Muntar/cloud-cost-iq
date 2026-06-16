package anomaly

import (
	"math"
)

func mean(values []float64,) float64 {

	var sum float64

	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values),)
}

func stddev(values []float64,) float64 {

	m := mean(values,)

	var total float64

	for _, v := range values {
	total +=
			math.Pow(v-m, 2,)
	}
	return math.Sqrt(total /float64(len(values),),
	)
}

func Detect(
	accountID string,
	service string,
	history []float64,
	current float64,
)(*Anomaly, bool,){

	if len(history,) < 7 {
		return nil,false
	}

	m := mean(history,)
	s := stddev(history,)

	threshold := m + (2 * s)

	if current < threshold {
		return nil, false
	}

	return &Anomaly{
		AccountID: accountID,
		Service: service,
		ExpectedCost: m,
		ActualCost: current,
		Deviation: current - m,
		Severity: highSeverity(current,m,),
	}, true
}

func highSeverity(current, expected float64,
) string {

	ratio := current / expected

	if ratio > 3 {
		return "critical"
	}

	if ratio > 2 {
		return "high"
	}

	return "medium"
}
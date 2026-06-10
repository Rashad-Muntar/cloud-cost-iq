package api

import (
	"net/http"
	"time"

	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/billing"
)

type Handlers struct {
	billingService *billing.Service
	aggregationService *aggregation.Service
}

func NewHandlers(billingService *billing.Service, aggregationService *aggregation.Service) *Handlers {
	return &Handlers{billingService: billingService, aggregationService: aggregationService}
}

func (h *Handlers) CreateCost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req billing.RecordCostInput
	err := h.billingService.RecordCost(
		ctx,
		req.AccountID,
		req.Service,
		req.Region,
		req.CostAmount,
		req.UsageAmount,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("cost recorded"))
}

func (h *Handlers) GetDailySummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dateStr := r.URL.Query().Get("date")
	accountID := r.URL.Query().Get("account_id")
	service := r.URL.Query().Get("service")
	date := time.Now()

	if dateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "invalid date format. use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		date = parsedDate
	}
	summary, err := h.billingService.GetDailySummary(ctx, date, accountID, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, summary)
}

func (h *Handlers) ExcuteDailySummaryJob(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	h.aggregationService.BuildDailySummaries(ctx, time.Now())
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Summary job initiated"))
}
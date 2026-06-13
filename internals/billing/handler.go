package billing

import (
	"net/http"
	"time"
	"github.com/cloud-cost-iq/internals/shared"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateCost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req RecordCostInput
	err := h.service.RecordCost(
		ctx,
		req.AwsAccountID,
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

func (h *Handler) GetDailySummary(w http.ResponseWriter, r *http.Request) {
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
	summary, err := h.service.GetDailySummary(ctx, date, accountID, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.WriteJSON(w, http.StatusOK, summary)
}
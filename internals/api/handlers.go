package api

import (
	"net/http"
	"time"
	"encoding/json"
	"github.com/cloud-cost-iq/internals/aggregation"
	"github.com/cloud-cost-iq/internals/billing"
	"github.com/cloud-cost-iq/internals/account"
)

type Handlers struct {
	billingService *billing.Service
	aggregationService *aggregation.Service
	accountService *account.Service
}

func NewHandlers(billingService *billing.Service, aggregationService *aggregation.Service) *Handlers {
	return &Handlers{billingService: billingService, aggregationService: aggregationService}
}

func (h *Handlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
    var input account.CreateAccountInput
    
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    acc, err := h.accountService.Create(r.Context(), input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    
    writeJSON(w, http.StatusCreated, acc)
}



func (h *Handlers) ListAccounts(w http.ResponseWriter, r *http.Request) {
    filter := account.AccountFilter{}
    
    accounts, err := h.accountService.List(r.Context(), filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    writeJSON(w, http.StatusOK, accounts)
}

func (h *Handlers) GetByAWSAccountID(w http.ResponseWriter, r *http.Request) {
    awsID := r.URL.Query().Get("aws_id")
    account, err := h.accountService.GetAccountByAWSID(r.Context(), awsID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    writeJSON(w, http.StatusOK, account)
}

func (h *Handlers) CreateCost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req billing.RecordCostInput
	err := h.billingService.RecordCost(
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
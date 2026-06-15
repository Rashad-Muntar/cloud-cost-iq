package analytics

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

func (h *Handler) Summary(
	w http.ResponseWriter,
	r *http.Request,
){
	
	accountID := r.URL.Query().Get("accountId")

	to := time.Now()

	from :=to.AddDate(0, 0, -30)

	result, err := h.service.GetCostSummary(r.Context(),
				CostQuery{
					AwsAccountID:accountID,
					From: from,
					To: to,
				},
			)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	shared.WriteJSON(w, http.StatusOK, result)
}

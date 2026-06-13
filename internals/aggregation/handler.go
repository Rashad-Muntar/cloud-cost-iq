package aggregation

import (
	"net/http"
	"time"
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

func (h *Handler) ExcuteDailySummaryJob(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	h.service.BuildDailySummaries(ctx, time.Now())
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Summary job initiated"))
}
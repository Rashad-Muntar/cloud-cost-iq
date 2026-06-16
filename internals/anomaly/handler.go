package anomaly

import (
	"net/http"
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

func (h *Handler) Run(w http.ResponseWriter,r *http.Request,){

	account := r.URL.Query().Get("accountId",)
	service := r.URL.Query().Get("service",)
	
	err := h.service.Run(r.Context(), account, service,)
	if err != nil {
		http.Error(w, err.Error(), 500,)
		return
	}

	w.WriteHeader(
		200,
	)
}
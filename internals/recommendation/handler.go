package recommendation

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

func (h *Handler) Generate(
	w http.ResponseWriter,
	r *http.Request,
){

	accountID := r.URL.Query().Get("accountId",)
	from := r.URL.Query().Get("from",)
	to := r.URL.Query().Get("to",)
	err := h.service.Generate(r.Context(),accountID, from, to,)
	if err != nil {
		http.Error(w, err.Error(), 500,)
		return
	}
	w.WriteHeader(200,)
}
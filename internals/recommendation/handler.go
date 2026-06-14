package recommendation

import (
	"net/http"
)


type Handler struct {
	service *Service
}

func (h *Handler) Generate(
	w http.ResponseWriter,
	r *http.Request,
){

	accountID :=
	r.URL.Query().
	Get(
		"accountId",
	)

	err :=
	h.service.
	Generate(
		r.Context(),
		accountID,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			500,
		)

		return
	}

	w.WriteHeader(
		200,
	)
}
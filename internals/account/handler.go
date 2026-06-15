package account

import (
	"fmt"
	"net/http"

	"encoding/json"

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

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handler hit")
    var input CreateAccountInput
    fmt.Println("Input from handler", input)
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    acc, err := h.service.Create(r.Context(), input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    
    shared.WriteJSON(w, http.StatusCreated, acc)
}

func (h *Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
    filter := AccountFilter{}
    
    accounts, err := h.service.List(r.Context(), filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    shared.WriteJSON(w, http.StatusOK, accounts)
}

func (h *Handler) GetByAWSAccountID(w http.ResponseWriter, r *http.Request) {
    awsID := r.URL.Query().Get("aws_id")
    account, err := h.service.GetAccountByAWSID(r.Context(), awsID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    shared.WriteJSON(w, http.StatusOK, account)
}
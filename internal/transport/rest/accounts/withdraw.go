package accounts

import (
	"bank/internal/services/account_bank"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func WithdrawHandler(service account_bank.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return
		}
		var req withdrawRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = service.Withdraw(r.Context(), id, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := withdrawResponse{
			Success: true,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

type withdrawRequest struct {
	Amount float64 `json:"amount"`
}

type withdrawResponse struct {
	Success bool `json:"success"`
}

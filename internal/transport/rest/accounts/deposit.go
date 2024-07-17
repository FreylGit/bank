package accounts

import (
	"bank/internal/services/account_bank"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func DepositHandler(service account_bank.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return
		}
		var req depositRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = service.Deposit(r.Context(), id, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := depositResponse{
			Success: true,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type depositRequest struct {
	Amount float64 `json:"amount"`
}

type depositResponse struct {
	Success bool `json:"success"`
}

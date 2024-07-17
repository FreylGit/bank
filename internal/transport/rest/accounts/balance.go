package accounts

import (
	"bank/internal/services/account_bank"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func BalanceHandler(service account_bank.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return
		}
		balance, err := service.GetBalance(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := balanceResponse{
			Balance: balance,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type balanceResponse struct {
	Balance float64 `json:"balance"`
}

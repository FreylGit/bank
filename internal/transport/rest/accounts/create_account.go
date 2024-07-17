package accounts

import (
	"bank/internal/services/account_bank"
	"encoding/json"
	"net/http"
)

func CreateAccountHandler(service account_bank.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := service.CreateAccount(r.Context())
		if err != nil {
			//Тут можно отлавливать разные ошибки и возвращать разные коды, но нужно поторопиться
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		response := createAccountResponse{
			Id: id,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type createAccountResponse struct {
	Id int `json:"id"`
}

package database

import "bank/internal/models/account"

type AccountBankRepository interface {
	CreateAccount() (int, error)
	GetAccount(id int) (account.Account, error)
	UpdateAccount(account.Account) error
}

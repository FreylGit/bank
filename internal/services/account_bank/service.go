package account_bank

import "context"

type Service interface {
	CreateAccount(ctx context.Context) (int, error)
	Deposit(ctx context.Context, id int, amount float64) error
	Withdraw(ctx context.Context, id int, amount float64) error
	GetBalance(ctx context.Context, id int) (float64, error)
}

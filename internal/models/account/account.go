package account

import (
	"errors"
	"github.com/govalues/decimal"
)

type Account struct {
	ID      int
	balance decimal.Decimal
}

func NewAccount(id int, balance decimal.Decimal) Account {
	return Account{
		ID:      id,
		balance: balance,
	}
}
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	amountDec, err := decimal.NewFromFloat64(amount)
	if err != nil {
		return err
	}
	a.balance, err = a.balance.Add(amountDec)

	return err
}
func (a *Account) Withdraw(amount float64) error {
	amountDec, err := decimal.NewFromFloat64(amount)
	if err != nil {
		return err
	}
	a.balance, err = a.balance.Sub(amountDec)
	return err
}
func (a Account) GetBalance() float64 {
	b, ok := a.balance.Float64()
	if !ok {
		return 0
	}

	return b
}

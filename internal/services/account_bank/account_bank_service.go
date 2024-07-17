package account_bank

import (
	"bank/internal/database"
	"bank/internal/models/account"
	"context"
	"errors"
	"log"
	"time"
)

type AccountBankService struct {
	repository database.AccountBankRepository
}

func NewService(repository database.AccountBankRepository) *AccountBankService {
	return &AccountBankService{
		repository: repository,
	}
}

func (a *AccountBankService) CreateAccount(ctx context.Context) (int, error) {
	resultChan := make(chan int)
	errorChan := make(chan error)
	defer close(resultChan)
	defer close(errorChan)
	go func() {
		id, err := a.repository.CreateAccount()
		log.Printf("Created Account ID: %d \t time:%s", id, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- id
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case err := <-errorChan:
		return 0, err
	case result := <-resultChan:
		return result, nil
	}
}

func (a *AccountBankService) Deposit(ctx context.Context, id int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	resultChan := make(chan *account.Account)
	errorChan := make(chan error)
	defer close(resultChan)
	defer close(errorChan)

	go func() {
		findAccount, err := a.repository.GetAccount(id)

		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- &findAccount
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errorChan:
		return err
	case findAccount := <-resultChan:
		err := findAccount.Deposit(amount)
		if err != nil {
			return err
		}
		log.Printf("Deposit Account ID: %d \t time:%s", id, time.Now().Format("2006-01-02 15:04:05"))

		go func() {
			err := a.repository.UpdateAccount(*findAccount)
			errorChan <- err
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errorChan:
			return err
		}
	}
}

func (a *AccountBankService) Withdraw(ctx context.Context, id int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	resultChan := make(chan *account.Account)
	errorChan := make(chan error)
	defer close(resultChan)
	defer close(errorChan)

	go func() {
		findAccount, err := a.repository.GetAccount(id)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- &findAccount
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errorChan:
		return err
	case findAccount := <-resultChan:
		if findAccount.GetBalance() < amount {
			return errors.New("insufficient balance to write off")
		}
		err := findAccount.Withdraw(amount)
		log.Printf("Withdraw Account ID: %d \t time:%s", id, time.Now().Format("2006-01-02 15:04:05"))

		if err != nil {
			return err
		}
		go func() {
			err := a.repository.UpdateAccount(*findAccount)
			errorChan <- err
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errorChan:
			return err
		}
	}
}

func (a *AccountBankService) GetBalance(ctx context.Context, id int) (float64, error) {
	resultChan := make(chan *account.Account)
	errorChan := make(chan error)
	defer close(resultChan)
	defer close(errorChan)

	go func() {
		findAccount, err := a.repository.GetAccount(id)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- &findAccount
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case err := <-errorChan:
		return 0, err
	case findAccount := <-resultChan:
		log.Printf("GetBalance Account ID: %d \t time:%s", id, time.Now().Format("2006-01-02 15:04:05"))
		return findAccount.GetBalance(), nil
	}
}

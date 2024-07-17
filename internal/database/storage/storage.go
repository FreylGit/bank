package storage

import (
	"bank/internal/database"
	"bank/internal/models/account"
	"errors"
	"github.com/govalues/decimal"
	"sync"
)

type Storage struct {
	mx       sync.Mutex
	accounts map[int]account.Account
	lastId   int
}

func New() database.AccountBankRepository {
	return &Storage{
		accounts: make(map[int]account.Account),
		lastId:   0,
	}
}

func (s *Storage) CreateAccount() (int, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.lastId == 0 {
		s.lastId = 1
	}
	newAccount := account.NewAccount(s.lastId, decimal.Zero)
	s.accounts[s.lastId] = newAccount
	s.lastId++
	return newAccount.ID, nil
}

func (s *Storage) GetAccount(id int) (account.Account, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	findAccount, ok := s.accounts[id]
	if !ok {
		return account.Account{}, errors.New("account not found")
	}

	return findAccount, nil
}

func (s *Storage) UpdateAccount(account account.Account) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	_, ok := s.accounts[account.ID]
	if !ok {
		return errors.New("account not found")
	}

	s.accounts[account.ID] = account
	return nil
}

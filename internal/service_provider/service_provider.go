package service_provider

import (
	"bank/internal/config"
	"bank/internal/database"
	"bank/internal/database/storage"
	"bank/internal/services/account_bank"
)

type ServiceProvider struct {
	config                *config.Config
	accountBankRepository database.AccountBankRepository
	accountBankService    account_bank.Service
}

func NewServiceProvider() ServiceProvider {
	return ServiceProvider{}
}

func (sp *ServiceProvider) Config() *config.Config {
	if sp.config == nil {
		config.LoadConfig(".env")
		sp.config = config.NewConfig()
	}

	return sp.config
}

func (sp *ServiceProvider) AccountBankRepository() database.AccountBankRepository {
	if sp.accountBankRepository == nil {
		sp.accountBankRepository = storage.New()
	}

	return sp.accountBankRepository
}
func (sp *ServiceProvider) AccountBankService() account_bank.Service {
	if sp.accountBankService == nil {
		sp.accountBankService = account_bank.NewService(sp.AccountBankRepository())
	}

	return sp.accountBankService
}

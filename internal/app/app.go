package app

import (
	"bank/internal/service_provider"
	"bank/internal/transport/rest/accounts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type App struct {
	sp service_provider.ServiceProvider
}

func NewApp() App {
	serviceProvider := service_provider.NewServiceProvider()
	app := App{
		sp: serviceProvider,
	}
	return app
}
func (a *App) Start() {
	r := a.initRouter()
	address := a.sp.Config().Host + ":" + a.sp.Config().Port
	err := http.ListenAndServe(address, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
}

func (a *App) initRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/accounts", accounts.CreateAccountHandler(a.sp.AccountBankService()))
	r.Post("/accounts/{id}/deposit", accounts.DepositHandler(a.sp.AccountBankService()))
	r.Post("/accounts/{id}/withdraw", accounts.WithdrawHandler(a.sp.AccountBankService()))
	r.Get("/accounts/{id}/balance", accounts.BalanceHandler(a.sp.AccountBankService()))
	return r
}

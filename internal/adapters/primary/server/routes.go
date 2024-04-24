package server

import (
	"net/http"

	"falcon-manager/internal/adapters/primary/server/handlers"
	"falcon-manager/internal/adapters/primary/server/middleware"
	"falcon-manager/internal/core/services"

	"github.com/gorilla/mux"
)

func initRoutes(accountService *services.AccountService) http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateAccountHandler(w, r, accountService)
	}).Methods(http.MethodPost)
	mux.HandleFunc("/account/{id:[0-9]+}", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAccountHandler(w, r, accountService)
	})).Methods(http.MethodGet)
	mux.HandleFunc("/account", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAccountsHandler(w, r, accountService)
	})).Methods(http.MethodGet)
	mux.HandleFunc("/account/{id:[0-9]+}", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.PutAccountHandler(w, r, accountService)
	})).Methods(http.MethodPut)
	mux.HandleFunc("/account/{id:[0-9]+}", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteAccountHandler(w, r, accountService)
	})).Methods(http.MethodDelete)
	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		handlers.Token(w, r, accountService)
	}).Methods(http.MethodPost)
	return mux
}

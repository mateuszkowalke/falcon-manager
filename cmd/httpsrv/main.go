package main

import (
	"log"

	"falcon-manager/internal/adapters/primary/server"
	"falcon-manager/internal/adapters/primary/server/config"
	pgdb "falcon-manager/internal/adapters/secondary"
	"falcon-manager/internal/core/services"
)

func main() {
	config.ConfigInit()

	store, err := pgdb.InitStore(config.DatabaseURL())
	if err != nil {
		log.Fatalf("error initializing store: %s\n", err)
	}
	defer store.Close()

	accountService := services.InitAccountService(store)

	server, err := server.InitServer(accountService)
	if err != nil {
		log.Fatalf("error initializing http server: %s\n", err)
	}

	err = server.Run()
	panic(err)
}

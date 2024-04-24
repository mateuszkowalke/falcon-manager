package config

import (
	"errors"
	"os"
)

type Config struct {
	init        bool
	jWTSecret   string
	databaseURL string
}

var globalConfig = Config{init: false}

var errConfigUninitialized = errors.New("config has to be initialized")

func ConfigInit() {
	globalConfig.init = true
	globalConfig.jWTSecret = getJWTSecret()
	globalConfig.databaseURL = getDatabaseURL()
}

func getJWTSecret() string {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		panic("JWT_SECRET missing")
	}
	return secret
}

func JWTSecret() string {
	if !globalConfig.init {
		panic(errConfigUninitialized)
	}
	return globalConfig.jWTSecret
}

func getDatabaseURL() string {
	secret, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		panic("DATABASE_URL missing")
	}
	return secret
}

func DatabaseURL() string {
	if !globalConfig.init {
		panic(errConfigUninitialized)
	}
	return globalConfig.databaseURL
}

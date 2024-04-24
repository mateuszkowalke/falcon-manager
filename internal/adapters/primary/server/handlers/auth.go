package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"falcon-manager/internal/adapters/primary/server/config"
	"falcon-manager/internal/core/services"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type TokenDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

func Token(w http.ResponseWriter, r *http.Request, accountService *services.AccountService) {
	secret := config.JWTSecret()
	var dto TokenDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	acc, err := accountService.GetAccountByEmail(dto.Username)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(dto.Password))
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := jwt.MapClaims{
		"iss": "falcon-manager",
		"sub": acc.Email,
		"exp": time.Hour * 24,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenResponse := TokenResponse{signedToken, "Bearer"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tokenResponse)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

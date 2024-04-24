package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"falcon-manager/internal/adapters/primary/server/middleware"
	"falcon-manager/internal/core/domain"
	"falcon-manager/internal/core/services"

	"github.com/gorilla/mux"
)

type CreateAccountDTO struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"userName,omitempty"`
	Firstname string `json:"firstName,omitempty"`
	Lastname  string `json:"lastName,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Sex       string `json:"sex,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
}

type PutAccountDTO struct {
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstName,omitempty"`
	Lastname  string `json:"lastName,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Sex       string `json:"sex,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
	Blocked   bool   `json:"blocked,omitempty"`
	Active    bool   `json:"active,omitempty"`
}

func statusFromDomainError(err error) int {
	switch err {
	case domain.ErrInvalidSexVariant:
		return http.StatusBadRequest
	case domain.ErrInvalidRoleVariant:
		return http.StatusBadRequest
	case domain.ErrInvalidEmail:
		return http.StatusBadRequest
	case domain.ErrMissingRequiredFields:
		return http.StatusBadRequest
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request, accountService *services.AccountService) {
	var dto CreateAccountDTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&dto)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	acc := domain.Account{
		Email:     dto.Email,
		Username:  dto.Username,
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Avatar:    dto.Avatar,
		Sex:       dto.Sex,
		Password:  dto.Password,
		Role:      dto.Role,
		Blocked:   false,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	validAcc, err := domain.NewAccount(&acc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(statusFromDomainError(err))
		return
	}
	safeAcc, err := accountService.CreateAccount(validAcc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(safeAcc)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAccountHandler(w http.ResponseWriter, r *http.Request, accountService *services.AccountService) {
	id := mux.Vars(r)["id"]
	validId, err := strconv.Atoi(id)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contextEmail, err := emailFromContext(r.Context())
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	acc, err := accountService.GetAccount(validId, contextEmail)
	if err != nil {
		// TODO
		// create not found error
		log.Println(err)
		w.WriteHeader(statusFromDomainError(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&acc)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAccountsHandler(w http.ResponseWriter, r *http.Request, accountService *services.AccountService) {
	contextEmail, err := emailFromContext(r.Context())
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accs, err := accountService.GetAccounts(contextEmail)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(statusFromDomainError(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&accs)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func PutAccountHandler(w http.ResponseWriter, r *http.Request, accountService *services.AccountService) {
	id := mux.Vars(r)["id"]
	validId, err := strconv.Atoi(id)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contextEmailRaw := r.Context().Value(middleware.ContextUserKey)
	contextEmail, ok := contextEmailRaw.(string)
	if !ok {
		// TODO
		log.Println("couldn't convert email context value to string")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	validContextEmail, err := domain.NewEmail(contextEmail)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	acc, err := accountService.GetAccount(validId, validContextEmail)
	if err != nil {
		// TODO
		// create not found error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var dto PutAccountDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if acc.Email != r.Context().Value(middleware.ContextUserKey) {
		// TODO
		// admin override
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	newAcc := domain.Account{
		Id:        validId,
		Username:  dto.Username,
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Avatar:    dto.Avatar,
		Sex:       dto.Sex,
		Password:  dto.Password,
		Role:      dto.Role,
		Blocked:   dto.Blocked,
		Active:    dto.Active,
	}
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	validAcc, err := domain.NewAccount(&newAcc)
	if err != nil {
		// TODO
		// create bad input error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	safeAcc, err := accountService.PutAccount(validAcc)
	if err != nil {
		// TODO
		// create bad input error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(safeAcc)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request, accountService *services.AccountService) {
	id := mux.Vars(r)["id"]
	validId, err := strconv.Atoi(id)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contextEmailRaw := r.Context().Value(middleware.ContextUserKey)
	contextEmail, ok := contextEmailRaw.(string)
	if !ok {
		// TODO
		log.Println("couldn't convert email context value to string")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	validContextEmail, err := domain.NewEmail(contextEmail)
	if err != nil {
		// TODO
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	acc, err := accountService.GetAccount(validId, validContextEmail)
	if err != nil {
		// TODO
		// create not found error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if acc.Email != r.Context().Value(middleware.ContextUserKey) {
		// TODO
		// admin override
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = accountService.DeleteAccount(validId)
	if err != nil {
		// TODO
		// create not found error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func emailFromContext(c context.Context) (domain.Email, error) {
	contextEmailRaw := c.Value(middleware.ContextUserKey)
	contextEmail, ok := contextEmailRaw.(string)
	if !ok {
		return "", errors.New("couldn't convert email context value to string")
	}
	validContextEmail, err := domain.NewEmail(contextEmail)
	if err != nil {
		return "", fmt.Errorf("couldn't create a valid email from string in context: %w", err)
	}
	return validContextEmail, nil
}

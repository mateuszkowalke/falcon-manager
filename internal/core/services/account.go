package services

import (
	"fmt"
	"time"

	"falcon-manager/internal/core/domain"
	"falcon-manager/internal/ports"
)

type AccountService struct {
	store ports.AccountStore
}

type AccountNotFound error

func InitAccountService(store ports.AccountStore) *AccountService {
	return &AccountService{store}
}

func (srv *AccountService) CreateAccount(acc *domain.ValidAccount) (*domain.SafeAccount, error) {
	err := srv.store.SaveAccount(acc)
	if err != nil {
		return nil, fmt.Errorf("error saving account: %w", err)
	}
	return acc.ToSafe(), nil
}

func (srv *AccountService) GetAccount(id int, contextUserEmail domain.Email) (*domain.SafeAccount, error) {
	contextAcc, err := srv.store.ReadAccountByEmail(string(contextUserEmail))
	if err != nil {
		return nil, fmt.Errorf("error getting account with email %s: %w", contextUserEmail, err)
	}
	if contextAcc.Id != id && contextAcc.Role != "admin" {
		return nil, domain.ErrUnauthorized
	}
	acc, err := srv.store.ReadAccount(id)
	if err != nil {
		return nil, fmt.Errorf("error getting account id %d: %w", id, err)
	}
	return acc.ToSafe(), nil
}

func (srv *AccountService) GetAccounts(contextUserEmail domain.Email) ([]*domain.SafeAccount, error) {
	contextAcc, err := srv.store.ReadAccountByEmail(string(contextUserEmail))
	if err != nil {
		return nil, fmt.Errorf("error getting account with email %s: %w", contextUserEmail, err)
	}
	if contextAcc.Role != "admin" {
		return nil, domain.ErrUnauthorized
	}
	accs, err := srv.store.ReadAccounts()
	if err != nil {
		return nil, fmt.Errorf("error getting accounts: %w", err)
	}
	safeAccs := make([]*domain.SafeAccount, 0)
	for _, acc := range accs {
		safeAccs = append(safeAccs, acc.ToSafe())
	}
	return safeAccs, nil
}

func (srv *AccountService) PutAccount(acc *domain.ValidAccount) (*domain.SafeAccount, error) {
	acc.UpdatedAt = time.Now()
	err := srv.store.UpdateAccount(acc)
	if err != nil {
		return nil, fmt.Errorf("error updating account with id %d: %w", acc.Id, err)
	}
	return acc.ToSafe(), nil
}

func (srv *AccountService) DeleteAccount(id int) error {
	err := srv.store.DeleteAccount(id)
	if err != nil {
		return fmt.Errorf("error updating account with id %d: %w", id, err)
	}
	return nil
}

func (srv *AccountService) GetAccountByEmail(email string) (*domain.ValidAccount, error) {
	acc, err := srv.store.ReadAccountByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error getting account with email %s: %w", email, err)
	}
	return acc, nil
}

package ports

import (
	"errors"

	"falcon-manager/internal/core/domain"
)

var ErrAccountNotFound = errors.New("account not found")

var ErrAccountDuplicate = errors.New("account with specified email already exists")

type AccountStore interface {
	// mutates passed ValidAccount by updating it's id
	SaveAccount(acc *domain.ValidAccount) error
	ReadAccount(id int) (*domain.ValidAccount, error)
	ReadAccountByEmail(email string) (*domain.ValidAccount, error)
	ReadAccounts() ([]*domain.ValidAccount, error)
	UpdateAccount(acc *domain.ValidAccount) error
	DeleteAccount(id int) error
}

// Package domain implements structs modelling the domain entities and value objects.
package domain

import (
	"errors"
	"fmt"
	"net/mail"
	"slices"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=./domain.go -package=mocks -destination=../../../mocks/mock_domain.go

type Email string

type Account struct {
	Id        int
	Email     string
	Username  string
	Firstname string
	Lastname  string
	Avatar    string
	Sex       string
	Password  string
	Role      string
	Blocked   bool
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ValidAccount struct {
	Id        int
	Email     Email
	Username  string
	Firstname string `db:"first_name"`
	Lastname  string `db:"last_name"`
	Avatar    string
	Sex       string
	Password  string
	Role      string
	Blocked   bool
	Active    bool
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type SafeAccount struct {
	Id        int    `json:"id"`
	Email     Email  `json:"email"`
	Username  string `json:"username"`
	Firstname string
	Lastname  string
	Avatar    string    `json:"avatar"`
	Sex       string    `json:"sex"`
	Role      string    `json:"role"`
	Blocked   bool      `json:"blocked"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var sexVals = []string{"MALE", "FEMALE", "OTHER"}

var roleVals = []string{"THERAPIST", "ADMIN", "USER"}

var ErrInvalidSexVariant = errors.New("'sex' value for account has to be one of 'MALE', 'FEMALE' or 'OTHER'")

var ErrInvalidRoleVariant = errors.New("'role' value for account has to be one of 'THERAPIST', 'ADMIN' or 'USER'")

var ErrInvalidEmail = errors.New("'email' field invalid")

var ErrMissingRequiredFields = errors.New("'email' or 'password' missing")

var ErrUnauthorized = errors.New("unauthorized")

func NewAccount(acc *Account) (*ValidAccount, error) {
	sexEnum := strings.ToUpper(acc.Sex)
	if !slices.Contains(sexVals, sexEnum) && sexEnum != "" {
		return &ValidAccount{}, ErrInvalidSexVariant
	}
	roleEnum := strings.ToUpper(acc.Role)
	if !slices.Contains(roleVals, roleEnum) && roleEnum != "" {
		return &ValidAccount{}, ErrInvalidRoleVariant
	}
	email, err := NewEmail(acc.Email)
	if err != nil {
		return &ValidAccount{}, err
	}
	if acc.Email == "" || acc.Password == "" {
		return &ValidAccount{}, ErrMissingRequiredFields
	}
	// default username is email
	if acc.Username == "" {
		acc.Username = acc.Email
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	return &ValidAccount{
		Id:        acc.Id,
		Email:     email,
		Username:  acc.Username,
		Firstname: acc.Firstname,
		Lastname:  acc.Lastname,
		Avatar:    acc.Avatar,
		Sex:       sexEnum,
		Password:  string(hashedPass),
		Role:      roleEnum,
		Blocked:   acc.Blocked,
		Active:    acc.Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (acc *ValidAccount) ToSafe() *SafeAccount {
	return &SafeAccount{
		Id:        acc.Id,
		Email:     acc.Email,
		Username:  acc.Username,
		Firstname: acc.Firstname,
		Lastname:  acc.Lastname,
		Avatar:    acc.Avatar,
		Sex:       acc.Sex,
		Role:      acc.Role,
		Blocked:   acc.Blocked,
		Active:    acc.Active,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}
}

func NewEmail(in string) (Email, error) {
	email, err := mail.ParseAddress(in)
	if err != nil {
		return "", ErrInvalidEmail
	}
	return Email(strings.Trim(email.String(), "<>")), nil
}

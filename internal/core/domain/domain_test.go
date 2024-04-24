package domain

import (
	"strings"
	"testing"
)

func TestNewAccount(t *testing.T) {
	t.Run("when called with valid account, then it should hash the password",
		func(t *testing.T) {
			acc := &Account{
				Id:        1,
				Email:     "test@example.com",
				Username:  "testusername",
				Firstname: "testfirstname",
				Lastname:  "testlastname",
				Avatar:    "http://avatar.com/test",
				Sex:       "MALE",
				Password:  "testpass",
				Role:      "USER",
				Blocked:   false,
				Active:    true,
			}
			validAcc, err := NewAccount(acc)
			if err != nil {
				t.Fatalf("err == %s, want%s\n", err, "nil")
			}
			if validAcc.Password == acc.Password {
				t.Errorf("validAcc.Password == %s, want %s\n", validAcc.Password, acc.Password)
			}
		})
	t.Run("when provided with valid sex, then it should uppercase it",
		func(t *testing.T) {
			acc := &Account{
				Id:        1,
				Email:     "test@example.com",
				Username:  "testusername",
				Firstname: "testfirstname",
				Lastname:  "testlastname",
				Avatar:    "http://avatar.com/test",
				Sex:       "male",
				Password:  "testpass",
				Role:      "USER",
				Blocked:   false,
				Active:    true,
			}
			validAcc, err := NewAccount(acc)
			if err != nil {
				t.Fatalf("err == %s, want%s\n", err, "nil")
			}
			if validAcc.Sex != strings.ToUpper(acc.Sex) {
				t.Errorf("validAcc.Sex == %s, want %s\n", validAcc.Sex, strings.ToUpper(acc.Sex))
			}
		})
	t.Run("when provided with valid role, then it should uppercase it",
		func(t *testing.T) {
			acc := &Account{
				Id:        1,
				Email:     "test@example.com",
				Username:  "testusername",
				Firstname: "testfirstname",
				Lastname:  "testlastname",
				Avatar:    "http://avatar.com/test",
				Sex:       "MALE",
				Password:  "testpass",
				Role:      "user",
				Blocked:   false,
				Active:    true,
			}
			validAcc, err := NewAccount(acc)
			if err != nil {
				t.Fatalf("Want err == nil, got err == %s\n", err)
			}
			if validAcc.Role != strings.ToUpper(acc.Role) {
				t.Errorf("validAcc.Role == %s, want %s\n", validAcc.Role, strings.ToUpper(acc.Role))
			}
		})
}

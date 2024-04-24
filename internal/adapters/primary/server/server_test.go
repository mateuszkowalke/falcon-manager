package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"falcon-manager/internal/adapters/primary/server/config"
	"falcon-manager/internal/adapters/primary/server/handlers"
	"falcon-manager/internal/core/domain"
	"falcon-manager/internal/core/services"
	"falcon-manager/internal/ports"
)

type MockAccountStore struct {
	lastId   int
	accounts []*domain.ValidAccount
}

func (store *MockAccountStore) SaveAccount(acc *domain.ValidAccount) error {
	store.lastId++
	acc.Id = store.lastId
	store.accounts = append(store.accounts, acc)
	return nil
}

func (store *MockAccountStore) ReadAccount(id int) (*domain.ValidAccount, error) {
	for _, acc := range store.accounts {
		if acc.Id == id {
			return acc, nil
		}
	}
	return nil, ports.ErrAccountNotFound
}

func (store *MockAccountStore) ReadAccountByEmail(email string) (*domain.ValidAccount, error) {
	for _, acc := range store.accounts {
		if string(acc.Email) == email {
			return acc, nil
		}
	}
	return nil, ports.ErrAccountNotFound
}

func (store *MockAccountStore) ReadAccounts() ([]*domain.ValidAccount, error) {
	return store.accounts, nil
}

func (store *MockAccountStore) UpdateAccount(acc *domain.ValidAccount) error {
	for i, dbAcc := range store.accounts {
		if dbAcc.Id == acc.Id {
			store.accounts[i] = acc
			return nil
		}
	}
	return ports.ErrAccountNotFound
}

func (store *MockAccountStore) DeleteAccount(id int) error {
	for i, acc := range store.accounts {
		if acc.Id == id {
			store.accounts = append(store.accounts[0:i], store.accounts[i+1:len(store.accounts)-1]...)
			return nil
		}
	}
	return ports.ErrAccountNotFound
}

var mockAccountStore = MockAccountStore{lastId: 0, accounts: []*domain.ValidAccount{}}

func (store *MockAccountStore) reset() {
	store.lastId = 0
	store.accounts = []*domain.ValidAccount{}
}

func router() http.Handler {
	accountService := services.InitAccountService(&mockAccountStore)
	srv, err := InitServer(accountService)
	if err != nil {
		panic(fmt.Sprintf("could not create test server: %s\n", err))
	}
	return srv.mux
}

func TestAccountCreateHandler(t *testing.T) {
	t.Run("when no body is sent, then it should return BadRequest status",
		func(t *testing.T) {
			body := []byte{}
			r := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusBadRequest
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when invalid body is sent, then it should return BadRequest status",
		func(t *testing.T) {
			body := "{\"key\":\"value\"}"
			r := httptest.NewRequest(http.MethodPost, "/account", strings.NewReader(body))
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusBadRequest
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when 'email' field is not a valid email address, then it should return BadRequest status",
		func(t *testing.T) {
			dto := handlers.CreateAccountDTO{Password: "test_password"}
			body, err := json.Marshal(dto)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusBadRequest
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when no 'email' field is sent, then it should return BadRequest status",
		func(t *testing.T) {
			dto := handlers.CreateAccountDTO{Password: "test_password"}
			body, err := json.Marshal(dto)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusBadRequest
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when no 'password' field is sent, then it should return BadRequest status",
		func(t *testing.T) {
			dto := handlers.CreateAccountDTO{Email: "test_email@example.com"}
			body, err := json.Marshal(dto)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusBadRequest
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when valid body is sent, then should return safe account data",
		func(t *testing.T) {
			dto := handlers.CreateAccountDTO{Email: "test_email@example.com", Password: "test_password"}
			body, err := json.Marshal(dto)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusCreated
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
			var res domain.SafeAccount
			decoder := json.NewDecoder(w.Result().Body)
			decoder.DisallowUnknownFields()
			err = decoder.Decode(&res)
			if err != nil {
				t.Errorf("err == %s, want %s\n", err, "nil")
			}
			// testing if username got the default value - email
			if res.Username != string(res.Email) {
				t.Errorf("username == %s, want %s\n", res.Username, res.Email)
			}
		})
}

func TestAccountGetHandler(t *testing.T) {
	os.Setenv("JWT_SECRET", "somesecret")
	os.Setenv("DATABASE_URL", "dburl")
	config.ConfigInit()
	t.Run("when trying to get user info by non-authenticated, then should return status unauthorized",
		func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/account/1", nil)
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusUnauthorized
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when trying to get info on current authenticated user, then should return it's data",
		func(t *testing.T) {
			router := router()
			mockAccountStore.reset()
			// create user
			userEmail := "test_email@example.com"
			userPass := "test_password"
			user, err := createUser(router, userEmail, userPass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// get token
			token, err := getAccessToken(router, userEmail, userPass)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/account/%d", user.Id), nil)
			r.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			expectedStatus := http.StatusOK
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
			var res domain.SafeAccount
			decoder := json.NewDecoder(w.Result().Body)
			decoder.DisallowUnknownFields()
			err = decoder.Decode(&res)
			if err != nil {
				panic(err)
			}
			// testing if username and email match
			if res.Email != user.Email {
				t.Errorf("email == %s, want %s\n", res.Email, user.Email)
			}
			if res.Username != string(user.Email) {
				t.Errorf("username == %s, want %s\n", res.Username, user.Email)
			}
		})
	t.Run("when trying to get info on another user by non-admin, then should return status unauthorized",
		func(t *testing.T) {
			router := router()
			mockAccountStore.reset()
			// create user 1
			userEmail := "test_email@example.com"
			userPass := "test_password"
			_, err := createUser(router, userEmail, userPass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// create user 2
			user2Email := "test_email2@example.com"
			user2Pass := "test_password"
			user2, err := createUser(router, user2Email, user2Pass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// get token
			token, err := getAccessToken(router, userEmail, userPass)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/account/%d", user2.Id), nil)
			r.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			expectedStatus := http.StatusUnauthorized
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
	t.Run("when trying to get info on another user by admin, then should return status OK and it's data",
		func(t *testing.T) {
			router := router()
			mockAccountStore.reset()
			// create user 1
			userEmail := "test_email@example.com"
			userPass := "test_password"
			_, err := createUser(router, userEmail, userPass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			mockAccountStore.accounts[0].Role = "admin"
			// create user 2
			user2Email := "test_email2@example.com"
			user2Pass := "test_password"
			user2, err := createUser(router, user2Email, user2Pass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// get token
			token, err := getAccessToken(router, userEmail, userPass)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/account/%d", user2.Id), nil)
			r.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			expectedStatus := http.StatusOK
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
			var res domain.SafeAccount
			decoder := json.NewDecoder(w.Result().Body)
			decoder.DisallowUnknownFields()
			err = decoder.Decode(&res)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// testing if username and email match
			if res.Email != user2.Email {
				t.Errorf("email == %s, want %s\n", res.Email, user2.Email)
			}
			if res.Username != string(user2.Username) {
				t.Errorf("username == %s, want %s\n", res.Username, user2.Username)
			}
		})
}

func TestAccountsGetHandler(t *testing.T) {
	os.Setenv("JWT_SECRET", "somesecret")
	os.Setenv("DATABASE_URL", "dburl")
	config.ConfigInit()
	t.Run("when trying to get all users info by non-authenticated, then should return status unauthorized",
		func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/account", nil)
			w := httptest.NewRecorder()
			router().ServeHTTP(w, r)
			expectedStatus := http.StatusUnauthorized
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})

	t.Run("when trying to get all users by admin, then should return status OK and data",
		func(t *testing.T) {
			router := router()
			mockAccountStore.reset()
			// create user 1
			userEmail := "test_email@example.com"
			userPass := "test_password"
			user, err := createUser(router, userEmail, userPass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			mockAccountStore.accounts[0].Role = "admin"
			// create user 2
			user2Email := "test_email2@example.com"
			user2Pass := "test_password"
			user2, err := createUser(router, user2Email, user2Pass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// get token
			token, err := getAccessToken(router, userEmail, userPass)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodGet, "/account", nil)
			r.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			expectedStatus := http.StatusOK
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
			var res []domain.SafeAccount
			decoder := json.NewDecoder(w.Result().Body)
			decoder.DisallowUnknownFields()
			err = decoder.Decode(&res)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// testing if username and email match
			if len(res) != 2 {
				t.Errorf("len(res) == %d, want %d\n", len(res), 2)
			}
			if res[0].Email != user.Email {
				t.Errorf("email == %s, want %s\n", res[0].Email, user.Email)
			}
			if res[0].Username != string(user.Username) {
				t.Errorf("username == %s, want %s\n", res[0].Username, user.Username)
			}
			if res[1].Email != user2.Email {
				t.Errorf("email == %s, want %s\n", res[1].Email, user2.Email)
			}
			if res[1].Username != string(user2.Username) {
				t.Errorf("username == %s, want %s\n", res[1].Username, user2.Username)
			}
		})

	t.Run("when trying to get all users by non-admin, then should return status unauthorized",
		func(t *testing.T) {
			router := router()
			mockAccountStore.reset()
			// create user 1
			userEmail := "test_email@example.com"
			userPass := "test_password"
			_, err := createUser(router, userEmail, userPass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// create user 2
			user2Email := "test_email2@example.com"
			user2Pass := "test_password"
			_, err = createUser(router, user2Email, user2Pass)
			if err != nil {
				t.Log(err)
				panic(err)
			}
			// get token
			token, err := getAccessToken(router, userEmail, userPass)
			if err != nil {
				panic(err)
			}
			r := httptest.NewRequest(http.MethodGet, "/account", nil)
			r.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			expectedStatus := http.StatusUnauthorized
			if w.Code != expectedStatus {
				t.Errorf("w.Code == %d, want %d\n", w.Code, expectedStatus)
			}
		})
}

func createUser(router http.Handler, email, password string) (domain.SafeAccount, error) {
	createUserDTO := handlers.CreateAccountDTO{Email: email, Password: password}
	body, err := json.Marshal(createUserDTO)
	if err != nil {
		return domain.SafeAccount{}, err
	}
	r := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	var createRes domain.SafeAccount
	decoder := json.NewDecoder(w.Result().Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&createRes)
	if err != nil {
		return domain.SafeAccount{}, err
	}
	return createRes, nil
}

func getAccessToken(router http.Handler, email, password string) (string, error) {
	dto := handlers.TokenDTO{Username: email, Password: password}
	body, err := json.Marshal(dto)
	if err != nil {
		panic(err)
	}
	r := httptest.NewRequest(http.MethodPost, "/token", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		return "", fmt.Errorf("error retrieving token for email \"%s\" and password \"%s\": %d", email, password, w.Code)
	}
	var tokenRes handlers.TokenResponse
	decoder := json.NewDecoder(w.Result().Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&tokenRes)
	if err != nil {
		return "", err
	}
	return tokenRes.AccessToken, nil
}

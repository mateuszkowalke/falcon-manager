package pgdb

import (
	"database/sql"
	"fmt"

	"falcon-manager/internal/core/domain"
	"falcon-manager/internal/ports"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgStore struct {
	db *sqlx.DB
}

func InitStore(connString string) (*PgStore, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return &PgStore{}, fmt.Errorf("error creating db connection: %w", err)
	}
	return &PgStore{
		db,
	}, nil
}

func (store *PgStore) Close() {
	store.db.Close()
}

func (store *PgStore) SaveAccount(acc *domain.ValidAccount) error {
	duplicateAcc, err := store.ReadAccountByEmail(string(acc.Email))
	if err != nil {
		return fmt.Errorf("error checking for duplicate email: %s", err)
	}
	if duplicateAcc != nil {
		return ports.ErrAccountDuplicate
	}
	rows, err := store.db.NamedQuery(`
    INSERT INTO
      account (
        email,
        username,
        first_name,
        last_name,
        avatar,
        sex,
        password,
        role,
        blocked,
        active,
        created_at,
        updated_at
      )
    VALUES
    (:email, :username, :firstname, :lastname, :avatar, :sex, :password, :role, :blocked, :active, :created_at, :updated_at)
    RETURNING id;
    `, acc)
	if err != nil {
		return fmt.Errorf("error creating account: %w", err)
	}
	var id int
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return fmt.Errorf("error reading inserted account's id: %w", err)
	}
	acc.Id = id
	return nil
}

func (store *PgStore) ReadAccount(id int) (*domain.ValidAccount, error) {
	row := store.db.QueryRowx(`
    SELECT
      id,
      email,
      username,
      first_name,
      last_name,
      avatar,
      sex,
      password,
      role,
      blocked,
      active,
      created_at,
      updated_at
    FROM
      account
    WHERE
      id = $1;
    `, id)
	var acc domain.ValidAccount
	err := row.StructScan(&acc)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrAccountNotFound
		}
		return nil, fmt.Errorf("error querying for account id %d: %w", id, err)
	}
	return &acc, nil
}

func (store *PgStore) ReadAccountByEmail(email string) (*domain.ValidAccount, error) {
	row := store.db.QueryRowx(`
    SELECT
      id,
      email,
      username,
      first_name,
      last_name,
      avatar,
      sex,
      password,
      role,
      blocked,
      active,
      created_at,
      updated_at
    FROM
      account
    WHERE
      email = $1;
    `, email)
	var acc domain.ValidAccount
	err := row.StructScan(&acc)
	if err != nil {
		return nil, fmt.Errorf("error querying for account with email %s: %w", email, err)
	}
	return &acc, nil
}

func (store *PgStore) ReadAccounts() ([]*domain.ValidAccount, error) {
	rows, err := store.db.Queryx(`
    SELECT
      id,
      email,
      username,
      first_name,
      last_name,
      avatar,
      sex,
      password,
      role,
      blocked,
      active,
      created_at,
      updated_at
    FROM
      account;
    `)
	if err != nil {
		return nil, fmt.Errorf("error querying for list of accounts: %w", err)
	}
	defer rows.Close()
	accs := make([]*domain.ValidAccount, 0)
	for rows.Next() {
		var acc domain.ValidAccount
		err = rows.StructScan(&acc)
		if err != nil {
			return nil, fmt.Errorf("error scanning row of accounts list: %w", err)
		}
		accs = append(accs, &acc)
	}
	return accs, nil
}

func (store *PgStore) UpdateAccount(acc *domain.ValidAccount) error {
	_, err := store.db.NamedExec(`
    UPDATE account SET
        username=:username,
        first_name=:firstname,
        last_name=:lastname,
        avatar=:avatar,
        sex=:sex,
        password=:password,
        role=:role,
        blocked=:blocked,
        active=:active,
        updated_at=:updated_at
    WHERE id=:id;
    `, acc)
	if err != nil {
		return fmt.Errorf("error updating account with id %d: %w", acc.Id, err)
	}
	return nil
}

func (store *PgStore) DeleteAccount(id int) error {
	_, err := store.db.Exec("DELETE FROM account WHERE id=$1;", id)
	if err != nil {
		return fmt.Errorf("error deleting account with id %d: %w", id, err)
	}
	return nil
}

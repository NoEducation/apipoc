package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(string) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(string) (*Account, error)
	GetAccountByEmail(string) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connectionString := "user=postgres password=mysecretpassword dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStorage) createAccountTable() error {
	command := `CREATE TABLE IF NOT EXISTS account (
		id VARCHAR(64) PRIMARY KEY,
		firstName VARCHAR(128) NOT NULL,
		lastName VARCHAR(128) NOT NULL,
		email VARCHAR(256) NOT NULL,
		password VARCHAR(256) NOT NULL,
		number BIGINT NOT NULL,
		balance NUMERIC NOT NULL,
		createdAt TIMESTAMP,

		CONSTRAINT ux_email_unique UNIQUE (email)
	);`

	_, err := s.db.Exec(command)

	if err != nil {
		log.Fatal("Error creating account table:", err)
	}

	return err
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	sql := `
		INSERT INTO account(id, firstName, lastName, email, password, number, balance, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := s.db.Exec(sql, account.Id,
		account.FirstName,
		account.LastName,
		account.Email,
		account.Password,
		account.Number,
		account.Balance,
		account.CreateAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", account)

	return nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	sql := `SELECT id, firstName, lastName, email, number, balance, createdAt FROM account`
	accounts, err := s.db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer accounts.Close()

	results := []*Account{}

	for accounts.Next() {
		account, err := s.scanIntoAccount(accounts)

		if err != nil {
			return nil, err
		}

		results = append(results, account)
	}

	return results, nil
}

func (s *PostgresStorage) DeleteAccount(id string) error {
	sql := "DELETE FROM account WHERE id = $1"

	_, err := s.db.Exec(sql, id)

	if err != nil {
		return fmt.Errorf("Error occurred while deleting account with id %s: %v", id, err)
	}

	return err
}

func (s *PostgresStorage) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountById(id string) (*Account, error) {
	sql := `SELECT id, firstName, lastName, email, password, number, balance, createdAt FROM account WHERE id = $1`
	account := s.db.QueryRow(sql, id)

	result := new(Account)
	err := account.Scan(&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.Password,
		&result.Number,
		&result.Balance,
		&result.CreateAt)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	return result, nil
}

func (s *PostgresStorage) GetAccountByEmail(email string) (*Account, error) {
	sql := `SELECT id, firstName, lastName, email, password, number, balance, createdAt FROM account WHERE email = $1`
	account := s.db.QueryRow(sql, email)

	result := new(Account)
	err := account.Scan(&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.Password,
		&result.Number,
		&result.Balance,
		&result.CreateAt)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	return result, nil
}

func (s *PostgresStorage) scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := &Account{}
	err := rows.Scan(
		&account.Id,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreateAt)

	if err != nil {
		return nil, err
	}

	return account, nil
}

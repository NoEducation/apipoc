package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(string) (*Account, error)
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
		number BIGINT NOT NULL,
		balance NUMERIC NOT NULL,
		createdAt TIMESTAMP
	);`

	_, err := s.db.Exec(command)

	if err != nil {
		log.Fatal("Error creating account table:", err)
	}

	return err
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	sql := `
		INSERT INTO account(id, firstName, lastName, number, balance, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := s.db.Exec(sql, account.Id,
		account.FirstName,
		account.LastName,
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
	sql := `SELECT id, firstName, lastName, number, balance, createdAt FROM account`
	accounts, err := s.db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer accounts.Close()

	results := []*Account{}

	for accounts.Next() {
		account := new(Account)
		err := accounts.Scan(&account.Id,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreateAt)

		if err != nil {
			return nil, err
		}

		results = append(results, account)
	}

	return results, nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountById(id string) (*Account, error) {
	sql := `SELECT id, firstName, lastName, number, balance, createdAt FROM account WHERE id = $1`
	account := s.db.QueryRow(sql, id)

	result := new(Account)
	err := account.Scan(&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Number,
		&result.Balance,
		&result.CreateAt)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	return result, nil
}

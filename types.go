package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type TransferAccount struct {
	ToAccountNumber int64   `json:"toAccountNumber"`
	Amount          float64 `json:"amount"`
}

type Account struct {
	Id        string    `json:"id"`
	FirstName string    `json:"fistName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreateAt  time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, password, email string) *Account {
	return &Account{
		Id:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Number:    int64(rand.Intn(1000000000)),
		Balance:   0.0,
		CreateAt:  time.Now().UTC(),
	}
}

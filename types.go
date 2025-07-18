package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TransferAccount struct {
	ToAccountNumber int64   `json:"toAccountNumber"`
	Amount          float64 `json:"amount"`
}

type Account struct {
	Id        string    `json:"id"`
	FirstName string    `json:"fistName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreateAt  time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000000)),
		Balance:   0.0,
		CreateAt:  time.Now().UTC(),
	}
}

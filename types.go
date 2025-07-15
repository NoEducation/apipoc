package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	Id        string  `json:"id"`
	FirstName string  `json:"fistName"`
	LastName  string  `json:"lastName"`
	Number    int64   `json:"number"`
	Balance   float64 `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000000)),
		Balance:   0.0,
	}
}

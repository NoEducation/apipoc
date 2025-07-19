package main

import (
	"fmt"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc := NewAccount("John", "Doe", "password123", "john.doe@example.com")

	fmt.Printf("%+v\n", acc)
}

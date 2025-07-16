package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStorage()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal("Error initializing storage:", err)
	}

	fmt.Printf("%+v\n", store)

	server := NewApiServer(":3000", store)
	server.Run()
}

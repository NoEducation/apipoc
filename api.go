package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddress string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader((status))
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewApiServer(listenAddress string) *ApiServer {
	return &ApiServer{listenAddress: listenAddress}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))

	log.Println("Registering handler, listening on", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	case http.MethodPut:
		return s.handleTransfer(w, r)
	}

	return fmt.Errorf("Unsupported method: %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("John", "Doe") // Example account creation

	WriteJson(w, http.StatusOK, account)

	return nil
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

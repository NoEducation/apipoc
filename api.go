package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	store         Storage
	listenAddress string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiErrorResponse struct {
	Error string `json:"error"`
}

type ApiSuccessResponse struct {
	Response map[string]any `json:"response"`
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader((status))
	return json.NewEncoder(w).Encode(v)
}

func NewApiServer(listenAddress string, store Storage) *ApiServer {
	return &ApiServer{listenAddress: listenAddress, store: store}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandleFunc(s.handleAccountById))
	router.HandleFunc("/account/transfer", makeHttpHandleFunc(s.handleTransferAccount))

	log.Println("Registering handler, listening on", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiErrorResponse{Error: err.Error()})
		}
	}
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccounts(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodPut:
		return s.handlePut(w, r)
	}

	return fmt.Errorf("Unsupported method: %s", r.Method)
}

func (s *ApiServer) handleAccountById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Unsupported method: %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)["id"]

	account, err := s.store.GetAccountById(id)

	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, account)

	return nil
}

func (s *ApiServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()

	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, accounts)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	message := ApiSuccessResponse{
		Response: map[string]any{
			"message": "Account created successfully",
			"account": account,
		},
	}

	return WriteJson(w, http.StatusOK, message)
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}

	message := ApiSuccessResponse{
		Response: map[string]any{
			"message":   "Account deleted successfully",
			"accountId": id,
		},
	}

	return WriteJson(w, http.StatusOK, message)
}

func (s *ApiServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	transferRequest := new(TransferAccount)

	if err := json.NewDecoder(r.Body).Decode(transferRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	return WriteJson(w, http.StatusOK, ApiSuccessResponse{
		Response: map[string]any{
			"message": "Account transferred successfully",
			"body":    transferRequest,
		},
	})
}

func (s *ApiServer) handlePut(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	listenAddr string
	store      Storage
}

func NewApiServer(listenAddr string, store Storage) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHanldeFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHanldeFunc(s.handleGetAccount))

	log.Println("JSON API server running on port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.hanldeCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	return fmt.Errorf("method not allowed: %v", r.Method)
}

func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccout("Reesav", "Gupta")
	id := mux.Vars(r)["id"]
	log.Printf("id: %v", id)
	return WriteJson(w, http.StatusOK, account)
}
func (s *APIserver) hanldeCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/*
-------------------------------------------------
********************UTILITIES********************
-------------------------------------------------
*/

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)               // always add status after the headers are set
	return json.NewEncoder(w).Encode(v) // json.NewEncoder takes in a io.writer and returns a Encoder, reposneWriter implements a io.writer
}

func makeHttpHanldeFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

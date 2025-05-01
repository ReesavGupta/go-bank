package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v) // json.NewEncoder takes in a io.writer and returns a Encoder, reposneWriter implements a io.writer
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHanldeFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error here
		}
	}

}

type APIserver struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", s.handleAccount)
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
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

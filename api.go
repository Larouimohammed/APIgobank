package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anthdm/Gobank/types"
	"github.com/gorilla/mux"
)



type APIServer struct {listenAddr string}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct{ Error string }

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/account",makeHTTPHandlefunc(s.handleAccount))
	router.HandleFunc("/account",makeHTTPHandlefunc(s.handleGetAccount))
	log.Println("JSON API server listeninnc(s.handleAccount))g on", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)

}
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == "PUT" {
		return s.handleTransfer(w, r)
	}
	return fmt.Errorf("unsupported method %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := types.NewAccount("aroui", "ee")
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}


func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
  	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlefunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})

		}
	}
}
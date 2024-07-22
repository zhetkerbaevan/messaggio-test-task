package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	db *sql.DB
	addr string
}

func NewAPIServer(db *sql.DB, addr string) *APIServer { //Create instance of APIServer
	return &APIServer{db: db, addr: addr}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter() 
	subrouter := router.PathPrefix("/api/v1").Subrouter()


	
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
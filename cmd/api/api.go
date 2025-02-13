package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/darshanparmar18/ecom/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)

	// to test wether it works without userhandler
	// router.HandleFunc("/login",handleLogin).Methods("POST")
	log.Println("listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}

// func handleLogin(w http.ResponseWriter, req *http.Request) {
// }

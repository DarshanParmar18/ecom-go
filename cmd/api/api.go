package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/darshanparmar18/ecom/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct{
	listenAddr string
	db *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
		db: db,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)	
	userHandler.RegisterRoutes(router)

	log.Println("listening on :",s.listenAddr)
	http.ListenAndServe(s.listenAddr, router) 
}



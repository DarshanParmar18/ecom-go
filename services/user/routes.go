package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(routes *mux.Router) {
	routes.HandleFunc("/login",h.handleLogin).Methods("POST")
	routes.HandleFunc("/resgister",h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, req *http.Request) {
	
}

func (h *Handler) handleRegister(w http.ResponseWriter, req *http.Request) {
	
}
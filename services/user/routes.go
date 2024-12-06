package user

import (
	"fmt"
	"net/http"

	"github.com/darshanparmar18/ecom/services/auth"
	"github.com/darshanparmar18/ecom/types"
	"github.com/darshanparmar18/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/users", h.getUsers).Methods("GET")

}

func (h *Handler) getUsers(w http.ResponseWriter, req *http.Request) {
	user, err := h.store.GetUsers()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
func (h *Handler) handleLogin(w http.ResponseWriter, req *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, req *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// check if the user exist
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", err))
		return
	}

	// convert the string password into hashed Password
	hashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// if user doesn't exist then create new user
	if err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashPassword,
	}); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "New User Registered Successfully!")
}

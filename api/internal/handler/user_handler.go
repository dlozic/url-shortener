package handler

import (
	"api/internal/model"
	"api/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		input model.User
		err   error
	)
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &model.User{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = h.userService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	user, err := h.userService.FindByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func SetupUserRoutes(router *mux.Router, userService service.UserService) {
	userHandler := NewUserHandler(userService)
	router.HandleFunc("/api/users", userHandler.Create).Methods("POST")
	router.HandleFunc("/api/users/find_by_email", userHandler.FindByEmail).Methods("GET")
}

package handlers

import (
	"encoding/json"
	"fmt"
	"gapi/pkg/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PostUserHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	var user UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := userService.CreateUser((*services.UserRequest)(&user)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user, err := userService.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	w.Header().Set("Content-Type", "application/json")

	user, err := userService.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	err = userService.DeleteUserById(id)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	var user UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := userService.UpdateUserById(id, (*services.UserRequest)(&user)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

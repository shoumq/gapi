package main

import (
	"fmt"
	"gapi/internal/handlers"
	"gapi/pkg/repositories"
	"gapi/pkg/services"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	repo := repositories.NewUserRepository()
	userService := services.NewUserService(repo)

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserHandler(w, r, userService)
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsersHandler(w, r, userService)
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostUserHandler(w, r, userService)
	}).Methods("POST")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUserHandler(w, r, userService)
	}).Methods("DELETE")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUserHandler(w, r, userService)
	}).Methods("PATCH")

	http.Handle("/", router)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

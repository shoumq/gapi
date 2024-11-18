package main

import (
	"fmt"
	"gapi/internal/handlers"
	"gapi/internal/middlewares"
	"gapi/pkg/repositories"
	"gapi/pkg/services"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	router.HandleFunc("/users/{id}/add_admin", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddAdminHandler(w, r, userService)
	}).Methods("POST")

	router.HandleFunc("/users/{id}/del_admin", func(w http.ResponseWriter, r *http.Request) {
		handlers.DelAdminHandler(w, r, userService)
	}).Methods("POST")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUserHandler(w, r, userService)
	}).Methods("DELETE")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUserHandler(w, r, userService)
	}).Methods("PATCH")

	router.Handle("/admin", middlewares.AdminMiddleware(http.HandlerFunc(handlers.AdminHandler)))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the home page!")
	})

	http.ListenAndServe(":8080", router)
}

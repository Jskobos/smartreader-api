package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllUsers : Returns a list of all users in the database.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// GetUser : Returns info about a specific user.
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// UpdateUser : Updates a User object.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteUser : Deletes a User record
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteMovieEndPoint : Deletes a movie from the database.
func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

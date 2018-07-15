package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jskobos/smartreader-api/db"
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

// GetAllTexts : Returns a list of all reading texts in the database.
func GetAllTexts(w http.ResponseWriter, r *http.Request) {
	db.BuildSchema()
	// fmt.Fprintln(w, "not implemented yet !")
}

// GetText : Retrieve a Text object by id.
func GetText(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// UpdateText : Update a Text object in the database.
func UpdateText(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteText : Deletes a Text record from the database.
func DeleteText(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", GetUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/text", GetAllTexts).Methods("GET")
	r.HandleFunc("/text/{id}", GetText).Methods("POST")
	r.HandleFunc("/text/{id}", UpdateText).Methods("PUT")
	r.HandleFunc("/text/{id}", DeleteText).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

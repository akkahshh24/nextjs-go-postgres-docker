package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func enableCORS() {

}

func jsonContentTypeMiddleware() {

}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/go-app/users", createUser(db)).Methods("POST")

	router.HandleFunc("/go-app/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/go-app/users/{id}", getUser(db)).Methods("GET")

	router.HandleFunc("go-app/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("go-app/users/{id}", deleteUser(db)).Methods("DELETE")

	// enable CORS and json parsing and unparsing middleware
	enhancedRouter := enableCORS(jsonContentTypeMiddleware(router))

	// start the server
	log.Fatal(http.ListenAndServe(":8080", enhancedRouter))
}

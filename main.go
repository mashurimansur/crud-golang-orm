package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/user", AllUseres).Methods("GET")
	myRouter.HandleFunc("/user/create", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/update/{id}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GO ORM Tutorial")

	InitialMigration()

	handleRequests()
}

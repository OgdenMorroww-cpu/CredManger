package main

import (
	"log"
	"net/http"

	"github.com/configs"
	"github.com/gorilla/mux"
	"github.com/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	configs.ConnectDatabase()

	router.HandleFunc("/credentials", handlers.GetCreds).Methods("GET")
	router.HandleFunc("/credentials", handlers.CreatCred).Methods("POST")
	router.HandleFunc("/credentials/{credentialID}", handlers.GetSingleCred).Methods("GET")
	router.HandleFunc("/credentials/{credentialID}", handlers.UpdateCred).Methods("PUT")
	router.HandleFunc("/credentials/{credentialsID}", handlers.DeleteCreds).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

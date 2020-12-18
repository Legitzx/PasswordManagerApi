package main

import (
	"./controller"
	//"./model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/update", controller.UpdateHandler).Methods("PUT")
	router.HandleFunc("/get", controller.GetVaultHandler).Methods("GET")

	log.Println("Server started and listening on http:.127.0.0.1:8080")

	http.ListenAndServe("127.0.0.1:8080", router)
}

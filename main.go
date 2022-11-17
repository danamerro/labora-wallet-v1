package main

import (
	"Labora_Wallet/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//ruta principal
	router.HandleFunc("/", controllers.Index)
	//ruta post
	router.HandleFunc("/search", controllers.CreateEventSearch).Methods("POST")
	//ruta get
	router.HandleFunc("/get", controllers.GetEvent).Methods("GET")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

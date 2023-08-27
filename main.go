package main

import (
	"go-rest-api/lets-go/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", api.HomeLink)
	router.HandleFunc("/event", api.CreateEvent).Methods("POST")
	router.HandleFunc("/events", api.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", api.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", api.UpdateEvent).Methods("PUT")
	router.HandleFunc("/events/{id}", api.DeleteEvent).Methods("DELETE")

	// log server status
	log.Println("Server is starting...")

	// Use the Gorilla Mux router
	http.Handle("/", router)

	// Start the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}

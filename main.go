package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"numbers-service/handlers"
)

func main() {
	router := mux.NewRouter()

	// Register endpoints
	router.HandleFunc("/odd", handlers.GetOddNumbers).Methods("GET")
	router.HandleFunc("/even", handlers.GetEvenNumbers).Methods("GET")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

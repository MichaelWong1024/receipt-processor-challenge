package main

import (
	"log" // For logging messages and errors
	"net/http" // For handling HTTP requests and responses
	"github.com/gorilla/mux" // For routing HTTP requests to handlers
	
	"receipt-processor/handlers"
)

func main() {
	r := mux.NewRouter() // Create a new router instance
	r.HandleFunc("/receipts/process", handlers.ProcessReceipts).Methods("POST")      // Register the processReceipts handler for POST requests at /receipts/process
	r.HandleFunc("/receipts/{id}/points", handlers.GetReceiptPoints).Methods("GET") // Register the getReceiptPoints handler for GET requests at /receipts/{id}/points

	log.Println("Server starting on port 8080...") // Log a message to indicate that the server is starting
	log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server on port 8080 and log any errors
}

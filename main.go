package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	"receipt-processor/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceipts).Methods("POST")      // Register the ProcessReceipts handler for POST requests at /receipts/process
	r.HandleFunc("/receipts/{id}/points", handlers.GetReceiptPoints).Methods("GET") // Register the GetReceiptPoints handler for GET requests at /receipts/{id}/points

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server on port 8080 and log any errors
}

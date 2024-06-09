package handlers

import (
	"encoding/json"
	"net/http" 
	"github.com/google/uuid" 
	"github.com/gorilla/mux"
	
	"receipt-processor/models" 
	"receipt-processor/utils" 
)

var ReceiptStore = make(map[string]models.Receipt)


// processReceipts handles POST requests to process receipts
func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt) // Decode the JSON request body into a Receipt object
	// Send a 400 Bad Request response if there was an error decoding the request body
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	receiptID := uuid.New().String() // Generate a unique ID for the receipt
	ReceiptStore[receiptID] = receipt // Store the receipt in the ReceiptStore map with the generated ID

	response := map[string]string{"id": receiptID} // Create a response map with the generated ID
	json.NewEncoder(w).Encode(response) // Encode the response map as JSON and send it in the response body
}


// getReceiptPoints handles GET requests to get points for a receipt
func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get the URL parameters from the request
	id, ok := vars["id"] // Get the "id" parameter from the URL
	// Send a 400 Bad Request response if the "id" parameter is missing
	if !ok {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest)
		return
	}

	receipt, ok := ReceiptStore[id] // Get the receipt from the ReceiptStore map using the ID
	// Send a 404 Not Found response if the receipt ID is not found in the store
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := utils.CalculatePoints(receipt)
	json.NewEncoder(w).Encode(map[string]int{"points": points}) // Encode the points as JSON and send it in the response body
}

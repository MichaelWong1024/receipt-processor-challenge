package main

import (
	"encoding/json" // For encoding and decoding JSON
	"log" // For logging messages and errors
	"net/http" // For handling HTTP requests and responses
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid" // For generating unique IDs
	"github.com/gorilla/mux" // For routing HTTP requests to handlers
)

// Receipt represents a receipt from a retailer
type Receipt struct {
	Retailer     string  `json:"retailer"`     // Name of the retailer
	PurchaseDate string  `json:"purchaseDate"` // Date of the purchase (YYYY-MM-DD)
	PurchaseTime string  `json:"purchaseTime"` // Time of the purchase (HH:MM)
	Items        []Item  `json:"items"`        // List of items on the receipt
	Total        float64 `json:"total"`        // Total amount paid
}

// Item represents an item on a receipt
type Item struct {
	ShortDescription string  `json:"shortDescription"` // Short description of the item
	Price            float64 `json:"price"`            // Price of the item
}

// ReceiptStore is a map of receipt IDs to Receipt objects as a database
var ReceiptStore = make(map[string]Receipt)

func main() {
	r := mux.NewRouter() // Create a new router instance
	r.HandleFunc("/receipts/process", processReceipts).Methods("POST")      // Register the processReceipts handler for POST requests at /receipts/process
	r.HandleFunc("/receipts/{id}/points", getReceiptPoints).Methods("GET") // Register the getReceiptPoints handler for GET requests at /receipts/{id}/points

	log.Println("Server starting on port 8080...") // Log a message to indicate that the server is starting
	log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server on port 8080 and log any errors
}

// processReceipts handles POST requests to process receipts
func processReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt) // Decode the JSON request body into a Receipt object
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Send a 400 Bad Request response if there was an error decoding the request body
		return
	}

	receiptID := uuid.New().String() // Generate a unique ID for the receipt
	ReceiptStore[receiptID] = receipt // Store the receipt in the ReceiptStore map with the generated ID

	response := map[string]string{"id": receiptID} // Create a response map with the generated ID
	json.NewEncoder(w).Encode(response) // Encode the response map as JSON and send it in the response body
}

// getReceiptPoints handles GET requests to get points for a receipt
func getReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get the URL parameters from the request
	id, ok := vars["id"] // Get the "id" parameter from the URL
	if !ok {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest) // Send a 400 Bad Request response if the "id" parameter is missing
		return
	}

	receipt, ok := ReceiptStore[id] // Get the receipt from the ReceiptStore map using the ID
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound) // Send a 404 Not Found response if the receipt ID is not found in the store
		return
	}

	points := calculatePoints(receipt) // Calculate points for the receipt
	json.NewEncoder(w).Encode(map[string]int{"points": points}) // Encode the points as JSON and send it in the response body
}

// calculatePoints calculates the points for a receipt
func calculatePoints(receipt Receipt) int {
	points := 0

	// Filter non-alphanumeric characters from the retailer name and count the remaining characters
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	alphanumericName := reg.ReplaceAllString(receipt.Retailer, "")
	points += len(alphanumericName)

	// 50 points if the total is a whole dollar amount
	if math.Mod(receipt.Total, 1.0) == 0 {
		points += 50
	}
	// 25 points if the total is a multiple of $0.25
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every 2 items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// 20% of the price for items with a short description that is a multiple of 3 characters
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	// 6 points if the purchase date is an odd day of the month
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the purchase time is between 2:00 PM and 4:00 PM
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

package models

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
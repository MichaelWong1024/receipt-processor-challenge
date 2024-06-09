package models

// a receipt from a retailer
type Receipt struct {
	Retailer     string  `json:"retailer"`     // retailer name
	PurchaseDate string  `json:"purchaseDate"` // purchase date (YYYY-MM-DD)
	PurchaseTime string  `json:"purchaseTime"` // purchase time (HH:MM)
	Items        []Item  `json:"items"`        // items on the receipt
	Total        float64 `json:"total"`        // total amount paid
}


// an item on a receipt
type Item struct {
	ShortDescription string  `json:"shortDescription"` // item description
	Price            float64 `json:"price"`            // item price
}
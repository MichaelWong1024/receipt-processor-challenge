package utils

import (
	"math"
	"regexp"
	"strings"
	"time"
	
	"receipt-processor/models"
)

// calculatePoints calculates the points for a receipt
func CalculatePoints(receipt models.Receipt) int {
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

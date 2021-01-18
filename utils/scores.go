package utils

import (
	"github.com/alexNgari/flooksTest/models"
)


// CalculateScore simulates a function to calculate creditscores of a person
func CalculateScore(borrower *models.Borrower) *models.CreditScore {
	// To be implemented
	return &models.CreditScore{
		BorrowerID: borrower.ID,
		CreditScore: 5,
	}
}

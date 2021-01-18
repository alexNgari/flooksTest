package models

import (
	"time"
)

// Borrower models a single customer with all the data required
type Borrower struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	AmmountBorrowed float64 `json:"ammountBorrowed"`
	AmmountPaidBack float64 `json:"ammountPaidBack"`
	DateSignedUp time.Time `json:"dateSignedUp"`
}


// CreditScore stores a user's credit score
type CreditScore struct {
	BorrowerID int64 `json:"borrowerId"`
	CreditScore float64 `json:"creditScore"`
}


// JSONData models the file with all borrowers info
type JSONData struct {
	Borrowers []Borrower `json:"borrowers"`
}


// CreditScores stores an array of users' credit scores
type CreditScores struct {
	CreditScores []CreditScore `json:"creditScores"`
}

package models

import "time"

type Transaction struct {
	ID          int       `json:"id"`          // Unique transaction ID
	WalletID    int       `json:"wallet_id"`   // Associated wallet ID
	Type        string    `json:"type"`        // Transaction type: topup, transfer
	Amount      float64   `json:"amount"`      // Amount of the transaction
	Timestamp   time.Time `json:"timestamp"`   // Time when the transaction occurred
	OtherParty  string    `json:"other_party"` // ID of the other wallet involved in the transfer (if applicable)
	Description string    `json:"description"` // Description
}

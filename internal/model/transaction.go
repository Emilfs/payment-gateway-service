package model

import "time"

type Transaction struct {
	ID            string    `json:"id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	CustomerID    string    `json:"customer_id"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
}

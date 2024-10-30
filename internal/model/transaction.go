package model

import (
	"time"
)

type Transaction struct {
	ID            string    `json:"id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	CustomerID    string    `json:"customer_id"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction_id"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
}

func (t *Transaction) GetByID(trxId string) error {
	// TODO: Implement this function
	return nil
}

func InsertTransaction(transaction Transaction) error {
	// TODO: Implement this function
	return nil
}

func (t *Transaction) UpdateStatus(trxId string, status string) error {
	// TODO: Implement this function
	return nil
}

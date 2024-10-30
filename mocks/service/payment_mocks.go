package service

import (
	"fmt"
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/model"
	"time"
)

// MockPaymentService is a mock of the PaymentService
type MockPaymentService struct {
	// Add fields to simulate internal state or expected behavior
	ShouldFail bool
}

// ProcessPayment mocks the processing of a payment request
func (m *MockPaymentService) ProcessPayment(request adapter.PaymentRequest) (model.Transaction, error) {
	if m.ShouldFail {
		return model.Transaction{}, fmt.Errorf("mock error processing payment")
	}
	return model.Transaction{
		ID:            "mock-transaction-id",
		Amount:        request.Amount,
		Currency:      request.Currency,
		CustomerID:    request.CustomerID,
		Status:        "processed",
		TransactionID: "mock-transaction-id-service",
		CreatedAt:     time.Now(),
	}, nil
}

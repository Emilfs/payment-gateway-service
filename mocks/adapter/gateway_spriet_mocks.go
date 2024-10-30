package adapter

import (
	"fmt"
	"payment-gateway-service/internal/adapter"
)

// MockGatewayB is a mock of the PaymentGateway interface for Gateway B
type MockGatewayB struct {
	// Add fields to simulate internal state or expected behavior
	ShouldFail bool
}

// ProcessPayment mocks the processing of a payment through Gateway B
func (m *MockGatewayB) ProcessPayment(request adapter.PaymentRequest) (adapter.PaymentResponse, error) {
	if m.ShouldFail {
		return adapter.PaymentResponse{}, fmt.Errorf("mock error processing payment")
	}
	return adapter.PaymentResponse{
		Success:       true,
		TransactionID: "mock-transaction-id-gateway-b",
		Message:       "Payment processed successfully by MockGatewayB",
	}, nil
}

package adapter

import (
	"fmt"
	"payment-gateway-service/internal/adapter"
)

// MockGatewayA is a mock of the PaymentGateway interface for Gateway A
type MockGatewayA struct {
	// Add fields to simulate internal state or expected behavior
	ShouldFail bool
}

// ProcessPayment mocks the processing of a payment through Gateway A
func (m *MockGatewayA) ProcessPayment(request adapter.PaymentRequest) (adapter.PaymentResponse, error) {
	if m.ShouldFail {
		return adapter.PaymentResponse{}, fmt.Errorf("mock error processing payment")
	}
	return adapter.PaymentResponse{
		Success:       true,
		TransactionID: "mock-transaction-id-gateway-a",
		Message:       "Payment processed successfully by MockGatewayA",
	}, nil
}

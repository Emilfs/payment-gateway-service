package adapter

type PaymentRequest struct {
	Amount     float64
	Currency   string
	CustomerID string
	Type       string
}

type CallbackRequest struct {
	TransactionID string
	Status        string
}

type PaymentResponse struct {
	Success       bool
	TransactionID string
	Message       string
}

// Extend the PaymentGateway interface
type PaymentGateway interface {
	ProcessPayment(request PaymentRequest) (PaymentResponse, error)
	ProcessWithdrawal(request PaymentRequest) (PaymentResponse, error)
	HandleCallback(request CallbackRequest) error
}

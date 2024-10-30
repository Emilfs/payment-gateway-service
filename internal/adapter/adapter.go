package adapter

type PaymentRequest struct {
	Amount     float64
	Currency   string
	CustomerID string
}

type PaymentResponse struct {
	Success       bool
	TransactionID string
	Message       string
}

type PaymentGateway interface {
	ProcessPayment(request PaymentRequest) (PaymentResponse, error)
}

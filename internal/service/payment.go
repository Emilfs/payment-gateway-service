package service

import (
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/model"
	"time"
)

type PaymentService struct {
	Gateway adapter.PaymentGateway
}

func NewPaymentService(gateway adapter.PaymentGateway) *PaymentService {
	return &PaymentService{Gateway: gateway}
}

func (s *PaymentService) ProcessPayment(request adapter.PaymentRequest) (model.Transaction, error) {
	response, err := s.Gateway.ProcessPayment(request)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction := model.Transaction{
		Amount:        request.Amount,
		Currency:      request.Currency,
		CustomerID:    request.CustomerID,
		Status:        "processed",
		TransactionID: response.TransactionID,
		CreatedAt:     time.Now(),
	}

	// TODO: Store transaction in DB

	return transaction, nil
}

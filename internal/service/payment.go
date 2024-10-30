package service

import (
	"fmt"
	"go.uber.org/zap"
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/model"
	"payment-gateway-service/pkg/constants"
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
		zap.L().Error("Failed to process payment", zap.Error(err), zap.String("CustomerID", request.CustomerID))
		return model.Transaction{}, err
	}

	transaction := model.Transaction{
		Amount:        request.Amount,
		Currency:      request.Currency,
		CustomerID:    request.CustomerID,
		Status:        constants.ProcessedTrxStatus,
		TransactionID: response.TransactionID,
		Type:          request.Type,
		CreatedAt:     time.Now(),
	}

	err = model.InsertTransaction(transaction)
	if err != nil {
		zap.L().Error("Failed to update transaction status", zap.Error(err), zap.String("TransactionID", transaction.TransactionID))
		return model.Transaction{}, fmt.Errorf("failed to update transaction status: %v", err)
	}
	zap.L().Info("Transaction stored", zap.String("TransactionID", transaction.TransactionID), zap.String("Status", transaction.Status))

	return transaction, nil
}

func (s *PaymentService) HandleCallback(request adapter.CallbackRequest) error {
	var transaction model.Transaction

	if transaction.ID == "" {
		zap.L().Error("Transaction not found for callback", zap.String("TransactionID", request.TransactionID))
		return fmt.Errorf("transaction not found")
	}

	err := transaction.UpdateStatus(request.TransactionID, request.Status)
	if err != nil {
		zap.L().Error("Failed to update transaction status", zap.Error(err), zap.String("TransactionID", transaction.TransactionID))
		return fmt.Errorf("failed to update transaction status: %v", err)
	}

	zap.L().Info("Transaction status updated", zap.String("TransactionID", transaction.TransactionID), zap.String("NewStatus", request.Status))
	return nil
}

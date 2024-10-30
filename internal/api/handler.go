package api

import (
	"net/http"
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	PaymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{PaymentService: paymentService}
}

func (h *PaymentHandler) HandlePaymentRequest(c *gin.Context) {
	var paymentRequest adapter.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		zap.L().Error("Failed to bind payment request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.PaymentService.ProcessPayment(paymentRequest)
	if err != nil {
		zap.L().Error("Failed to process payment", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	zap.L().Info("Payment processed successfully", zap.String("TransactionID", transaction.TransactionID))
	c.JSON(http.StatusOK, transaction)
}

func (h *PaymentHandler) HandleCallback(c *gin.Context) {
	var callbackRequest adapter.CallbackRequest
	if err := c.ShouldBindJSON(&callbackRequest); err != nil {
		zap.L().Error("Failed to bind callback request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PaymentService.HandleCallback(callbackRequest)
	if err != nil {
		zap.L().Error("Failed to process callback", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process callback"})
		return
	}

	zap.L().Info("Callback processed successfully", zap.String("TransactionID", callbackRequest.TransactionID))
	c.JSON(http.StatusOK, gin.H{"message": "Callback processed successfully"})
}

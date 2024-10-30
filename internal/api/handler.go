package api

import (
	"net/http"
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/service"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.PaymentService.ProcessPayment(paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *PaymentHandler) HandleCallback(c *gin.Context) {
	var callbackRequest adapter.CallbackRequest
	if err := c.ShouldBindJSON(&callbackRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PaymentService.HandleCallback(callbackRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process callback"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Callback processed successfully"})
}

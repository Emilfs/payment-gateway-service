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

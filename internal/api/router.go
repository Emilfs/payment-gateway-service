package api

import (
	"payment-gateway-service/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(paymentService *service.PaymentService) *gin.Engine {
	router := gin.Default()
	paymentHandler := NewPaymentHandler(paymentService)

	router.POST("/payment", paymentHandler.HandlePaymentRequest)
	router.POST("/callback", paymentHandler.HandleCallback)

	return router
}

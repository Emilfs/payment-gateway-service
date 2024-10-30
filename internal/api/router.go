package api

import (
	"payment-gateway-service/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns a new Gin router with all routes configured
func SetupRouter(paymentService *service.PaymentService) *gin.Engine {
	router := gin.Default()
	paymentHandler := NewPaymentHandler(paymentService)

	// Define all your routes here
	router.POST("/payment", paymentHandler.HandlePaymentRequest)

	// Future routes can be added here as well
	// router.GET("/another-endpoint", anotherHandlerFunction)

	return router
}

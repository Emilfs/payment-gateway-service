package main

import (
	"fmt"
	"log"
	"net/http"
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/api"
	"payment-gateway-service/internal/service"
	"payment-gateway-service/pkg/config"
	"payment-gateway-service/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig("./pkg/config") // point to config files location
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	_, err = logger.InitLogger()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	gatewayA := &adapter.GatewayPaybro{APIEndpoint: "http://example.com/gatewayA"}
	//gatewayB := &adapter.GatewayB{APIEndpoint: "http://example.com/gatewayB"}

	var chosenGateway adapter.PaymentGateway
	chosenGateway = gatewayA // TODO: Implement so that chosenGateway switch depending on the request

	paymentService := service.NewPaymentService(chosenGateway)

	router := gin.Default()
	paymentHandler := api.NewPaymentHandler(paymentService)

	router.POST("/payment", paymentHandler.HandlePaymentRequest)

	serverPort := cfg.ServerPort
	if serverPort == "" {
		serverPort = "8080" // Default port if not specified in configuration
	}
	zap.L().Info("Starting server...", zap.String("port", serverPort))
	if err := router.Run(fmt.Sprintf(":%s", serverPort)); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}

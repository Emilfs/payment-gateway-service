package adapter

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GatewayPaybro struct {
	APIEndpoint string
}

func (g *GatewayPaybro) ProcessPayment(request PaymentRequest) (PaymentResponse, error) {
	requestData, err := json.Marshal(request)
	if err != nil {
		return PaymentResponse{}, err
	}

	resp, err := http.Post(g.APIEndpoint, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return PaymentResponse{}, err
	}
	defer resp.Body.Close()

	var response PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return PaymentResponse{}, err
	}

	return response, nil
}

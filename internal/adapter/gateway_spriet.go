package adapter

import (
	"bytes"
	"net/http"
)

type GatewaySpriet struct {
	APIEndpoint string
}

func (g *GatewaySpriet) ProcessPayment(request PaymentRequest) (PaymentResponse, error) {
	// TODO: Handle XML data read
	requestData := "<xml>...</xml>"

	client := &http.Client{}
	req, err := http.NewRequest("POST", g.APIEndpoint, bytes.NewBufferString(requestData))
	if err != nil {
		return PaymentResponse{}, err
	}
	req.Header.Add("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		return PaymentResponse{}, err
	}
	defer resp.Body.Close()

	// TODO: Convert XML response to PaymentResponse
	var response PaymentResponse

	return response, nil
}

func (g *GatewaySpriet) HandleCallback(request CallbackRequest) error {
	// TODO: Consider if we want this or not
	return nil
}

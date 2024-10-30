package adapter

import (
	"encoding/json"
)

type GatewayPaybro struct {
	APIEndpoint string
}

func (g *GatewayPaybro) ProcessPayment(request PaymentRequest) (PaymentResponse, error) {
	_, err := json.Marshal(request)
	if err != nil {
		return PaymentResponse{}, err
	}

	//resp, err := http.Post(g.APIEndpoint, "application/json", bytes.NewBuffer(requestData))
	resp, err := PaymentResponse{
		Success:       true,
		TransactionID: "1234",
		Message:       "SUCCESS",
	}, nil // TODO: use mock here and remove this line, until then these codes will be commented out
	if err != nil {
		return PaymentResponse{}, err
	}
	//defer resp.Body.Close()

	//var response PaymentResponse
	//if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
	//	return PaymentResponse{}, err
	//}

	return resp, nil
}

func (g *GatewayPaybro) HandleCallback(request CallbackRequest) error {
	// TODO: Consider if we want want this func or not
	return nil
}

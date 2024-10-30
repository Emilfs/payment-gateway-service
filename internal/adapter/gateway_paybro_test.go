package adapter

import (
	"payment-gateway-service/pkg/constants"
	"reflect"
	"testing"
)

func TestGatewayPaybro_ProcessPayment(t *testing.T) {
	type fields struct {
		APIEndpoint string
	}
	type args struct {
		request PaymentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    PaymentResponse
		wantErr bool
	}{
		{
			name: "SUCCESS",
			fields: fields{
				APIEndpoint: "",
			},
			args: args{
				request: PaymentRequest{
					Amount:     100,
					Currency:   "AED",
					CustomerID: "1234",
					Type:       constants.DepositPayment,
				},
			},
			want: PaymentResponse{
				Success:       true,
				TransactionID: "1234",
				Message:       "SUCCESS",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GatewayPaybro{
				APIEndpoint: tt.fields.APIEndpoint,
			}
			got, err := g.ProcessPayment(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayPaybro.ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GatewayPaybro.ProcessPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGatewayPaybro_HandleCallback(t *testing.T) {
	type fields struct {
		APIEndpoint string
	}
	type args struct {
		request CallbackRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GatewayPaybro{
				APIEndpoint: tt.fields.APIEndpoint,
			}
			if err := g.HandleCallback(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("GatewayPaybro.HandleCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

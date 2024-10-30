package service

import (
	"payment-gateway-service/internal/adapter"
	"payment-gateway-service/internal/model"
	"reflect"
	"testing"
)

func TestNewPaymentService(t *testing.T) {
	type args struct {
		gateway adapter.PaymentGateway
	}
	tests := []struct {
		name string
		args args
		want *PaymentService
	}{
		{
			name: "SUCCESS",
			args: args{
				gateway: &adapter.GatewayPaybro{},
			},
			want: &PaymentService{Gateway: &adapter.GatewayPaybro{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentService(tt.args.gateway); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_ProcessPayment(t *testing.T) {
	type fields struct {
		Gateway adapter.PaymentGateway
	}
	type args struct {
		request adapter.PaymentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PaymentService{
				Gateway: tt.fields.Gateway,
			}
			got, err := s.ProcessPayment(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentService.ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.ProcessPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_HandleCallback(t *testing.T) {
	type fields struct {
		Gateway adapter.PaymentGateway
	}
	type args struct {
		request adapter.CallbackRequest
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
			s := &PaymentService{
				Gateway: tt.fields.Gateway,
			}
			if err := s.HandleCallback(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("PaymentService.HandleCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

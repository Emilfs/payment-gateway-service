package api

import (
	"payment-gateway-service/internal/service"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewPaymentHandler(t *testing.T) {
	type args struct {
		paymentService *service.PaymentService
	}
	tests := []struct {
		name string
		args args
		want *PaymentHandler
	}{
		{
			name: "SUCCESS",
			args: args{
				paymentService: &service.PaymentService{},
			},
			want: &PaymentHandler{PaymentService: &service.PaymentService{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentHandler(tt.args.paymentService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentHandler_HandlePaymentRequest(t *testing.T) {
	type fields struct {
		PaymentService *service.PaymentService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PaymentHandler{
				PaymentService: tt.fields.PaymentService,
			}
			h.HandlePaymentRequest(tt.args.c)
		})
	}
}

func TestPaymentHandler_HandleCallback(t *testing.T) {
	type fields struct {
		PaymentService *service.PaymentService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PaymentHandler{
				PaymentService: tt.fields.PaymentService,
			}
			h.HandleCallback(tt.args.c)
		})
	}
}

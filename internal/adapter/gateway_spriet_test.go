package adapter

import (
	"reflect"
	"testing"
)

func TestGatewaySpriet_ProcessPayment(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GatewaySpriet{
				APIEndpoint: tt.fields.APIEndpoint,
			}
			got, err := g.ProcessPayment(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewaySpriet.ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GatewaySpriet.ProcessPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGatewaySpriet_HandleCallback(t *testing.T) {
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
			g := &GatewaySpriet{
				APIEndpoint: tt.fields.APIEndpoint,
			}
			if err := g.HandleCallback(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("GatewaySpriet.HandleCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package model

import (
	"payment-gateway-service/pkg/constants"
	"testing"
	"time"
)

func TestGetByID(t *testing.T) {
	type args struct {
		trxId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS",
			args: args{
				trxId: "1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetByID(tt.args.trxId); (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertTransaction(t *testing.T) {
	type args struct {
		transaction Transaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS",
			args: args{
				transaction: Transaction{
					ID:            "2345",
					Amount:        1000,
					Currency:      "IDR",
					CustomerID:    "2345",
					Status:        constants.FailedTrxStatus,
					TransactionID: "2345",
					Type:          constants.DepositPayment,
					CreatedAt:     time.Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertTransaction(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("InsertTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_UpdateStatus(t *testing.T) {
	type fields struct {
		ID            string
		Amount        float64
		Currency      string
		CustomerID    string
		Status        string
		TransactionID string
		Type          string
		CreatedAt     time.Time
	}
	type args struct {
		trxId  string
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS",
			fields: fields{
				ID:            "1234",
				Amount:        100,
				Currency:      "AED",
				CustomerID:    "1234",
				Status:        constants.ProcessedTrxStatus,
				TransactionID: "1234",
				Type:          constants.WithdrawalPayment,
				CreatedAt:     time.Now(),
			},
			args: args{
				trxId:  "1234",
				status: constants.CompletedTrxStatus,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				ID:            tt.fields.ID,
				Amount:        tt.fields.Amount,
				Currency:      tt.fields.Currency,
				CustomerID:    tt.fields.CustomerID,
				Status:        tt.fields.Status,
				TransactionID: tt.fields.TransactionID,
				Type:          tt.fields.Type,
				CreatedAt:     tt.fields.CreatedAt,
			}
			if err := tr.UpdateStatus(tt.args.trxId, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

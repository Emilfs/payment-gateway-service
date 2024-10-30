package logger

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name    string
		want    *zap.Logger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitLogger()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

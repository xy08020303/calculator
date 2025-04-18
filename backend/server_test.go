package server

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	calculatorv1 "github.com/xy08020303/calculator/backend/gen/go/proto/calculator/v1"
)

func TestCalculatorServer(t *testing.T) {
	server := &CalculatorServer{}
	ctx := context.Background()

	tests := []struct {
		name    string
		req     *calculatorv1.CalculationRequest
		want    float64
		wantErr bool
		errCode connect.Code
	}{
		{
			name:    "addition",
			req:     &calculatorv1.CalculationRequest{Num1: 5, Num2: 3, Operation: "+"},
			want:    8,
			wantErr: false,
		},
		{
			name:    "subtraction",
			req:     &calculatorv1.CalculationRequest{Num1: 5, Num2: 3, Operation: "-"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "multiplication",
			req:     &calculatorv1.CalculationRequest{Num1: 5, Num2: 3, Operation: "*"},
			want:    15,
			wantErr: false,
		},
		{
			name:    "division",
			req:     &calculatorv1.CalculationRequest{Num1: 6, Num2: 3, Operation: "/"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "division by zero",
			req:     &calculatorv1.CalculationRequest{Num1: 6, Num2: 0, Operation: "/"},
			wantErr: true,
			errCode: connect.CodeInvalidArgument,
		},
		{
			name:    "unknown operation",
			req:     &calculatorv1.CalculationRequest{Num1: 6, Num2: 3, Operation: "x"},
			wantErr: true,
			errCode: connect.CodeInvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := connect.NewRequest(tt.req)
			res, err := server.Calculate(ctx, req)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if connect.CodeOf(err) != tt.errCode {
					t.Errorf("expected error code %v, got %v", tt.errCode, connect.CodeOf(err))
				}
				return
			}

			if err != nil {
				t.Fatalf("Calculate() error = %v", err)
			}
			if res.Msg.Result != tt.want {
				t.Errorf("Calculate() got = %v, want %v", res.Msg.Result, tt.want)
			}
		})
	}
}

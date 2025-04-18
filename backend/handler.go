package main

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
)

type CalculatorService struct{}

func (s *CalculatorService) Calculate(
	ctx context.Context,
	req *connect.Request[calculatorv1.CalculationRequest],
) (*connect.Response[calculatorv1.CalculationResponse], error) {
	a := req.Msg.A
	b := req.Msg.B
	op := req.Msg.Op
	var result float64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("division by zero"))
		}
		result = a / b
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported operator"))
	}

	return connect.NewResponse(&calculatorv1.CalculationResponse{Result: result}), nil
}

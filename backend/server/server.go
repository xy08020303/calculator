package server

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	calculatorv1 "github.com/xy08020303/calculator/backend/gen/go/calculator/v1"
)

type CalculatorServer struct{}

func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[calculatorv1.CalculationRequest],
) (*connect.Response[calculatorv1.CalculationResponse], error) {
	num1 := req.Msg.Num1
	num2 := req.Msg.Num2
	op := req.Msg.Operation

	var result float64
	var err error

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			err = errors.New("division by zero")
		} else {
			result = num1 / num2
		}
	default:
		err = fmt.Errorf("unknown operation: %s", op)
	}

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&calculatorv1.CalculationResponse{
		Result: result,
	})
	return res, nil
}

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	calcv1 "github.com/xy08020303/calculator-app/backend/gen/go/calculator/v1"
	calcv1connect "github.com/xy08020303/calculator-app/backend/gen/go/calculator/v1/calculatorv1connect"
)

type CalculatorService struct{}

func (s *CalculatorService) Add(ctx context.Context, req *connect.Request[calcv1.BinaryOpRequest]) (*connect.Response[calcv1.BinaryOpResponse], error) {
	return connect.NewResponse(&calcv1.BinaryOpResponse{Result: req.Msg.A + req.Msg.B}), nil
}
func (s *CalculatorService) Sub(ctx context.Context, req *connect.Request[calcv1.BinaryOpRequest]) (*connect.Response[calcv1.BinaryOpResponse], error) {
	return connect.NewResponse(&calcv1.BinaryOpResponse{Result: req.Msg.A - req.Msg.B}), nil
}
func (s *CalculatorService) Mul(ctx context.Context, req *connect.Request[calcv1.BinaryOpRequest]) (*connect.Response[calcv1.BinaryOpResponse], error) {
	return connect.NewResponse(&calcv1.BinaryOpResponse{Result: req.Msg.A * req.Msg.B}), nil
}
func (s *CalculatorService) Div(ctx context.Context, req *connect.Request[calcv1.BinaryOpRequest]) (*connect.Response[calcv1.BinaryOpResponse], error) {
	if req.Msg.B == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}
	return connect.NewResponse(&calcv1.BinaryOpResponse{Result: req.Msg.A / req.Msg.B}), nil
}

func main() {
	mux := http.NewServeMux()
	path, handler := calcv1connect.NewCalculatorServiceHandler(&CalculatorService{})
	mux.Handle(path, handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

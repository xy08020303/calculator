package main

import (
	"log"
	"net/http"
	calcv1connect "https://github.com/xy08020303/calculator-app/backend/gen/go/calculator/calculatorconnect"
	calcv1 "https://github.com/xy08020303/calculator-app/backend/gen/go/calculator"
)

func main() {
	mux := http.NewServeMux()

	service := &CalculatorService{}
	path, handler := calcv1.NewCalculatorServiceHandler(service)
	mux.Handle(path, handler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}

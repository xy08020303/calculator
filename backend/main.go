package backend

import (
	"log"
	"net/http"

	"github.com/xy08020303/calculator/backend/gen/go/proto/calculator/v1/calculatorv1connect"
	"github.com/xy08020303/calculator/backend/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	calculator := &server.CalculatorServer{}
	mux := http.NewServeMux()
	path, handler := calculatorv1connect.NewCalculatorServiceHandler(calculator)
	mux.Handle(path, handler)
	log.Println("Server started on :8080")
	http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

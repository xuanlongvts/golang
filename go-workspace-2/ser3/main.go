package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	const serverAddr string = "127.0.0.1:8083"

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Service 3 ping")
		_, _ = w.Write([]byte("3 Service"))
	})
	server := http.Server{
		Addr:    serverAddr,
		Handler: router,
	}
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		_, _ = fmt.Fprintf(os.Stderr, "HTTP server error %v\n", err)
	}
}
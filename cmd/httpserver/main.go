package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"sathwik.work/internal/request"
	"sathwik.work/internal/response"
	"sathwik.work/internal/server"
)

const port = 42069

func main() {
	s, err := server.Serve(port, func(w io.Writer, req *request.Request) *server.HandlerError {
		if req.RequestLine.RequestTarget == "/problemclient" {
			return &server.HandlerError{
				StatusCode: response.StatusBadRequest,
				Message:    "Error on the client side\n",
			}
		} else if req.RequestLine.RequestTarget == "/problemserver" {
			return &server.HandlerError{
				StatusCode: response.StatusInternalServerError,
				Message:    "Error on the server side\n",
			}
		} else {
			return &server.HandlerError{
				StatusCode: response.StatusOK,
				Message:    "Success\n",
			}
		}
	})

	if err != nil {
		log.Fatalf("Error stating server: %v", err)
	}
	defer s.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}

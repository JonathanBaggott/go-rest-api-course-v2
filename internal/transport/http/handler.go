package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type CommentService interface{}

// Handler is a struct that handles HTTP requests.
type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

// NewHandler creates a new instance of the Handler struct with the provided CommentService.
func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	// Create a new mux.Router instance.
	h.Router = mux.NewRouter()
	// Map the routes to their respective handlers.
	h.mapRoutes()

	// Create a new http.Server instance and assign it to the Handler's Server field.
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

// mapRoutes defines the routes and their corresponding handlers.
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

// Serve starts the HTTP server and handles graceful shutdown.
func (h *Handler) Serve() error {
	// Start the HTTP server in a goroutine.
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	// Create a channel to receive OS interrupt signals.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Wait for the interrupt signal to be received on the channel 'c'.
	<-c

	// Create a context with a timeout to gracefully shutdown the server.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil
}

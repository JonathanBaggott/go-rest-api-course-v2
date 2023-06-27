package http

import (
	"fmt"
	"net/http"

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

// Serve starts the HTTP server and listens for incoming requests.
func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

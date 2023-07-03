package http

import (
	"context"
	"encoding/json"
	"net/http"

	"log"

	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/comment"
	"github.com/gorilla/mux"
)

// CommentService defines the interface for comment operations
type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

// Response represents the response structure
type Response struct {
	Message string
}

// PostComment handles the HTTP POST request for creating a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment

	// Decode the request body into a Comment struct
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	// Call the PostComment method of the CommentService to create a new comment
	cmt, err := h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Print(err)
		return
	}

	// Encode the comment as JSON and send it in the response
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

// GetComment handles the HTTP GET request for retrieving a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the comment ID is provided
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the GetComment method of the CommentService to retrieve the comment by ID
	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encode the comment as JSON and send it in the response
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// UpdateComment handles the HTTP PUT request for updating a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the comment ID is provided
	if id == "" {
		// Set the HTTP response status code to indicate a 400 Bad Request.
		// This status code indicates that the server cannot process the client's request due to a client error.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment

	// Decode the request body into a Comment struct
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	// Update the comment by ID using the UpdateComment method of the CommentService
	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Print(err)
		// Set the HTTP response status code to indicate a 500 Internal Server Error.
		// This status indicates that an unexpected error occurred on the server side.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encode the updated comment as JSON and send it in the response
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// DeleteComment handles the HTTP DELETE request for deleting a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the comment ID is provided
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the DeleteComment method of the CommentService to delete the comment by ID
	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encode a success message as JSON and send it in the response
	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}

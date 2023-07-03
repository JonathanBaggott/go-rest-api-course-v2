package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods that our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service - is the struct on which all our logic will be built
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service (kind of like a constructor method)
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment retrieves a comment by ID
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retreiving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// UpdateComment updates a comment by ID.
// It invokes the Store interface's UpdateComment method to update the comment in the data store.
func (s *Service) UpdateComment(
	ctx context.Context,
	ID string,
	updatedCmt Comment,
) (Comment, error) {
	// The returned Comment object is assigned to the cmt variable, and the error (if any) is assigned to the err variable
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println("error updating comment")
		// Returns an empty Comment object along with the received error
		return Comment{}, err
	}
	// Return the updated Comment object and a nil error if there are no errors, indicating a successful update
	return cmt, nil
}

// DeleteComment deletes a comment by ID
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	// Call the DeleteComment method of the Store interface to delete the comment by ID
	return s.Store.DeleteComment(ctx, id)
}

// PostComment creates a new comment
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	// Call the PostComment method of the Store interface to create a new comment
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}

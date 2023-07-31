//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/comment"

	"github.com/stretchr/testify/assert"
)

// TestCommentDatabase is a test function for testing comment-related database operations.
func TestCommentDatabase(t *testing.T) {
	// t.Run represents a sub-test within TestCommentDatabase.

	// Sub-test to test creating a comment.
	t.Run("test create comment", func(t *testing.T) {
		// Create a new database instance.
		db, err := NewDatabase()
		// Assert that there is no error in creating the database.
		assert.NoError(t, err)

		// Post a new comment to the database and get the created comment.
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		// Assert that there is no error in posting the comment.
		assert.NoError(t, err)

		// Retrieve the newly created comment from the database using its ID.
		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		// Assert that the slug of the retrieved comment is equal to the one posted.
		assert.Equal(t, "slug", newCmt.Slug)
	})

	// Sub-test to test deleting a comment.
	t.Run("test delete comment", func(t *testing.T) {
		// Create a new database instance.
		db, err := NewDatabase()
		assert.NoError(t, err)

		// Post a new comment to the database and get the created comment.
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "jono",
			Body:   "body",
		})
		assert.NoError(t, err)

		// Delete the comment from the database using its ID.
		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		// Attempt to retrieve the deleted comment from the database using its ID.
		_, err = db.GetComment(context.Background(), cmt.ID)
		// Assert that there is an error in retrieving the deleted comment.
		assert.Error(t, err)
	})
}

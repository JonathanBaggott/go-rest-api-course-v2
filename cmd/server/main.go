package main

import (
	"fmt"

	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/comment"
	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/db"
	transportHttp "github.com/JonathanBaggott/go-rest-api-course-v2/internal/transport/http"
)

// Run - responsible for the instantiation and startup of our Go application
func Run() error {
	fmt.Println("starting up our application")

	// Connect to the database
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database!")
		return err
	}

	// Migrate the database schema
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
	}

	// Create a new comment service instance and inject the database
	cmtService := comment.NewService(db)

	// Create an HTTP handler and inject the comment service
	httpHandler := transportHttp.NewHandler(cmtService)

	// Start the HTTP server and handle requests
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	fmt.Println("successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}

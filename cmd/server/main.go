package main

import (
	"fmt"

	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/comment"
	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/db"
	transportHttp "github.com/JonathanBaggott/go-rest-api-course-v2/internal/transport/http"
)

// Run - is going to be responsible for the instantiation and startup of our Go application

func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database!")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
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

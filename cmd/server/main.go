package main

import (
	"context"
	"fmt"

	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/comment"
	"github.com/JonathanBaggott/go-rest-api-course-v2/internal/db"
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

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
			Slug:   "manual-test",
			Author: "Jono",
			Body:   "Hello World",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
	))

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

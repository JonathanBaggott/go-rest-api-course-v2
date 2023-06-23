package main

import (
	"fmt"

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

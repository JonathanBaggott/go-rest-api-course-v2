package main

import "fmt"

// Run - is going to be responsible for the instantiation and startup of our Go application

func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}

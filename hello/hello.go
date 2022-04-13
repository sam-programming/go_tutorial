package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger
	log.SetPrefix("greetings: ") //prefix the log/error message
	log.SetFlags(0)

	message, err := greetings.Hello("Sam")
	if err != nil {
		log.Fatal(err) // log.Fatal() kills the program
	}

	fmt.Println(message)
}

package main

import (
	"fmt"
	"log"

	"get-set-go.com/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	message, error := greetings.Hello("sagar")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}

package main

import (
	"fmt"

	"get-set-go.com/greetings"
)

func main() {
	message := greetings.Hello("Sagar")
	fmt.Println(message)
}

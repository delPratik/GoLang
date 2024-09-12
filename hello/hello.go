package main

import (
	"fmt"
	"module/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Pratik")
	fmt.Println(message)
}

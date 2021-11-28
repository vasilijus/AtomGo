package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message, errors := greetings.Hello("Sergio")
	// message, errors := greetings.Hello("")
	if errors == nil {
		fmt.Println(message)
	} else {
		fmt.Println(errors)
	}

}

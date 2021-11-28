package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Sergio")
	fmt.Println(message)
}

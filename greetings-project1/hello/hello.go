package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Ashik")
	fmt.Println(message)
}

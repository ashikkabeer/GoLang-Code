package greetings

import "fmt"

//thisfunctionreturnsastring
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

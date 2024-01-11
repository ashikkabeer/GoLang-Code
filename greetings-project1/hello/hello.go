package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

//functions with argument expects no error

//	func main() {
//		message := greetings.Hello("Ashik")
//		fmt.Println(message)
//	}
func main() {
	//set properties for the pre-defined logger
	//the log entry prifix and flag to disable printing the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	//request a greeting message
	//message, err := greetings.Hello("Ashik")
	//if err => print the console and exit
	if err != nil {
		log.Fatal(err)
		//	Fatal is equivalent to Print() followed by a call to os.Exit(1).
		// so it print the error and stops the process
		//	https://pkg.go.dev/log#Fatal
	}

	//if no err => print the message
	fmt.Println(messages)
}

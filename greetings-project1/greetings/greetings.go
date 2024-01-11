package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// function return greetings for the person named
// this function returns a string

//func Hello(name string) string {
//	message := fmt.Sprintf("Hi, %v. Welcome!", name)
//	return message;
//}

// here the return data is a string and an error
// if error. return empty string with error
// if not error er
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos function returns a map that associates each of the named people with greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate a name with each greeting
	//initialize a map with the following syntax: make(map[key-type]value-type).
	messages := make(map[string]string)
	//loop through  the received slice of names,
	//call the hello function to get the message for each name
	//In Go, you can't have unused variables, and
	//using _ allows you to discard values that you receive from functions or range iterations when
	//you're not going to use them. It helps to keep the code clean and avoids the need to create a variable
	//that you won't use.
	for _, name := range names {
		//range returns two values: the index of the current item in the loop and a copy of the item's value.
		//You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
		//range names iterates over each element in the names slice.
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map associate the retrieved message with name
		messages[name] = message

	}
	return messages, nil
}

// return a set of greeting messages
func randomFormat() string {
	//a slice of message formats
	formats := []string{
		"Hi %v. Welcome~",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	//return a string by index
	return formats[rand.Intn(len(formats))]
}

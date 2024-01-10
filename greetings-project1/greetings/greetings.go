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

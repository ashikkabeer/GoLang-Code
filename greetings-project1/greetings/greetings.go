package greetings

import (
	"errors"
	"fmt"
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
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

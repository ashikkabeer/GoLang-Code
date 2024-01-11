package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Ashik"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Ashik")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ashik") = %q, %v, want match for %#q, nil`, msg, err, want)

	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Ahmed"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)

	if err != nil || !want.MatchString(message) {
		t.Errorf("Hello(%q) = %q, %v, want %q, nil", name, message, err, want)
	}
}


// test empty name
func TestHelloEmptyName(t *testing.T) {
	message, err := Hello("")

	if message != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, message, err)
    }
}
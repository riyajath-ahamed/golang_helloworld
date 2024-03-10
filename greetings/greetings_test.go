package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Riyajath"
	game := "Dota 2"
	want := regexp.MustCompile(`\b` + name + `\b.*\b` + game + `\b`)
	msg, err := Hello("Riyajath", "Dota 2")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Riyajath", "Dota 2") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("", "Dota 2")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("", "Dota 2") = %q, %v, want "", error`, msg, err)
	}
}

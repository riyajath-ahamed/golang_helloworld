package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string, game string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name, game)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name, "Dota 2")
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {

	formats := []string{
		"Hi, %v. Welcome %v!",
		"Great %v to see you, %v!",
		"Hail, %v! Well met %v!",
	}

	return formats[rand.Intn(len(formats))]
}

package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is empty")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Static Greeting! Hello there %s", name)
	return message, nil
}

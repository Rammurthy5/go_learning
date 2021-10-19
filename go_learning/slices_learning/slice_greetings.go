package slices_learning

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func SliceGreet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is empty")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func SliceGreetMultiPeople(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := SliceGreet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {

	formats := []string{
		"Hi %v, welcome!",
		"great to see you %v",
		"Hail %v!, Glad you are here!!",
	}
	return formats[rand.Intn(len(formats))]
}

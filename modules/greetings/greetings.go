package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Note the := operator is a shortcut for declaring and initializing a variable in one line, but this is only
	// available inside a function. Go uses the value on the right to determine the variable's type. The long version of
	// this code is:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map (collection of key/value pairs) to associate names with messages.
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomGreeting() string {
	greetings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return greetings[rand.Intn(len(greetings))]
}

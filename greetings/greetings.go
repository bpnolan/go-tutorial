package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of ta set of greeting messages. The returned 
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

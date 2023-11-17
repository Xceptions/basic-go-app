package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
// In go, a func whose name starts with a capital
// letter can be called by another function not in
// that package. It is known as Exported names
func Hello(name string) (string, error) {

	//If no name was given, return an error with a message
	if name == "" {
		return "No name was given", errors.New("empty name")
	}
	
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


// randomFormat returns one of a set of greeting messages. The
// returned message is selected at random
func randomFormat() string {

	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
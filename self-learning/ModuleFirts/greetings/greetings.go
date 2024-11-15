package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Hello %v! How are you today?",
		"Salut %v! Comment allez-vous?",
		"Bonjour %v! Comment puis-je vous aider?",
		"Hola %v! ¿Cómo estás?",
	}
	return formats[rand.Intn(len(formats))]
}

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
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

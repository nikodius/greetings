package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// funcion devuleve saludo
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat()) // falla test
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

func randomFormat() string {
	formats := []string{
		"Hola, %v Bienvenido!",
		"Que bueno verte, %v",
		"Saludo, %v encantado de conocerte",
		"saludos mortal %v",
	}
	return formats[rand.Intn(len(formats))]
}

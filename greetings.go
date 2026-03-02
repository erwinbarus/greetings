package greetings

import ( 
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", erros.New("empty name")
	}

	return fmt.Sprintf("Hi %s. Welcome!!", name), nil
}
package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == ""{
	return "", errors.New("Empty name")
	}
	message := fmt.Sprintf("Hi, %v. Wecome!", name)
	return message, nil
}


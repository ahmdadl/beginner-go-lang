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

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomSentence(), name)
    return message, nil
}

func randomSentence() string {
    sentences := []string {
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    index := rand.Intn(len(sentences))

    // fmt.Println(index)

    return sentences[index]
}
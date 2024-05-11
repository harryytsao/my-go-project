package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// cretes prefix string and returns zero or whatever it's set to
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("jeff", ""))
}

/*
Test Driven Development (TDD) ⇒  Write a failing test and see it fail so we know we
have written a relevant test for our requirements and seen that it produces an easy
to understand description of the failure ⇒ Write the smallest amount of code to
pass the case ⇒ Refactor to ensure well-crafted code
*/

package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return getGreetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world!", "Umba-yumba"))
}

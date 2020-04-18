package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const defaultName = "World"

// Hello : It says hello
func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return prefixGreeting(language) + name
}

func prefixGreeting(language string) (prefix string) {
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

func main() {
	fmt.Println(Hello("world", ""))
	// fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

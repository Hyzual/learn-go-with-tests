package main

import (
	"fmt"
)

const spanishParameter = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const defaultName = "World"

// Hello : It says hello
func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	if language == spanishParameter {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
	// fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

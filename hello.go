package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

// Hello : It says hello
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
	// fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

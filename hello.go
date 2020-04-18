package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
	// fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

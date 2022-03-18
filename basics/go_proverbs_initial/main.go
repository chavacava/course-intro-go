package main

import (
	"fmt"
)

const location = "Remote"
const proverb = "Don't panic"

var name string

func main() {
	name = "Salva"
	from := `Jujuy`
	var courseDurationHs int = 10

	fmt.Printf("Hello, fellow %s Gophers!\n", location)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("We'll learn GO the next %d hours!", courseDurationHs)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
}

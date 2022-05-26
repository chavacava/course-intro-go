package main

import (
	"fmt"
)

var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {
	fmt.Printf("There are %d Go proverbs:\n", len(proverbs))

	count := countChars(proverbs)

	fmt.Printf("The total of characters of the proverbs is: %d\n", count)
}

func countChars(sentences []string) int {
	total := make(chan int)
	partial := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < len(sentences); i++ {
			sum += <-partial
		}
		total <- sum
	}()

	for _, p := range sentences {
		go func(p string) {
			partial <- len(p)
		}(p)
	}

	return <-total
}

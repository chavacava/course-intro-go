package main

import (
	"fmt"
)

func main() {
	n := 10
	fmt.Printf("There are 19 Go proverbs\n")
	fmt.Printf("Here's one: %s\n", getProverb(n))
}

// getProverb returns the n'th Go proverb
func getProverb(n int) string {
	p := proverbs[n]
	return fmt.Sprintf("#%d %s", n, p)
}

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

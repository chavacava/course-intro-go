package main

import (
	"fmt"
	"strings"
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

	upCassedProverbs := toUpperCase(proverbs)
	fmt.Printf("%v\n", upCassedProverbs)
	fmt.Printf("%v\n", len(upCassedProverbs))

}

const workersCount = 19

func toUpperCase(sentences []string) []string {
	jobs := make(chan string, 5)
	// producer
	go func() {
		for _, s := range sentences {
			jobs <- s
		}
		close(jobs)
	}()

	done := make(chan bool)
	jobResults := make(chan string, 5)
	for i := 0; i < workersCount; i++ {
		// worker
		go func() {
			for job := range jobs {
				jobResults <- strings.ToUpper(job)
			}
			done <- true
		}()
	}

	result := []string{}
	for {
		select {
		case upCassed := <-jobResults:
			result = append(result, upCassed)
		case <-done:
			return result
		}
	}

	return result
}

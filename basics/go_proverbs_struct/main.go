package main

import (
	"fmt"
)

type proverb struct {
	text string
	url  string
}

var proverbs = []proverb{
	{
		text: "Don't communicate by sharing memory, share memory by communicating.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s",
	},
	{
		text: "Concurrency is not parallelism.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=3m42s",
	},
	{
		text: "Channels orchestrate; mutexes serialize.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=4m20s",
	},
	{
		text: "The bigger the interface, the weaker the abstraction.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s",
	},
	{
		text: "Make the zero value useful.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s",
	},
	{
		text: "interface{} says nothing.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s",
	},
	{
		text: "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s",
	},
	{
		text: "A little copying is better than a little dependency.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s",
	},
	{
		text: "Syscall must always be guarded with build tags.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m10s",
	},
	{
		text: "Cgo must always be guarded with build tags.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m53s",
	},
	{
		text: "Cgo is not Go.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=12m37s",
	},
	{
		text: "With the unsafe package there are no guarantees.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=13m49s",
	},
	{
		text: "Clear is better than clever.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s",
	},
	{
		text: "Reflection is never clear.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=15m22s",
	},
	{
		text: "Errors are values.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=16m13s",
	},
	{
		text: "Don't just check errors, handle them gracefully.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=17m25s",
	},
	{
		text: "Design the architecture, name the components, document the details.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=18m09s",
	},
	{
		text: "Documentation is for users.",
		url:  "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=19m07s",
	},
	{
		text: "Don't panic.",
		url:  "https://github.com/golang/go/wiki/CodeReviewComments#dont-panic",
	},
}

func main() {
	p := proverbs[0]
	fmt.Printf("There are 19 Go proverbs\n")
	fmt.Printf("Here's one: %s\n", p.text)
	fmt.Printf("You can watch Rob's explanation here: %s\n", p.url)
}

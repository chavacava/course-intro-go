package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getAll)
	http.ListenAndServe(":8080", nil)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	for _, p := range proverbs {
		fmt.Fprintf(w, "#%d\t%s\n", p.number, p.text)
	}
}

type proverb struct {
	number int
	text   string
	url    string
}

var proverbs = map[int]proverb{
	1: {
		number: 1,
		text:   "Don't communicate by sharing memory, share memory by communicating.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s",
	},
	2: {
		number: 2,
		text:   "Concurrency is not parallelism.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=3m42s",
	},
	3: {
		number: 3,
		text:   "Channels orchestrate; mutexes serialize.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=4m20s",
	},
	4: {
		number: 4,
		text:   "The bigger the interface, the weaker the abstraction.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s",
	},
	5: {
		number: 5,
		text:   "Make the zero value useful.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s",
	},
	6: {
		number: 6,
		text:   "interface{} says nothing.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s",
	},
	7: {
		number: 7,
		text:   "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s",
	},
	8: {
		number: 8,
		text:   "A little copying is better than a little dependency.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s",
	},
	9: {
		number: 9,
		text:   "Syscall must always be guarded with build tags.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m10s",
	},
	10: {
		number: 10,
		text:   "Cgo must always be guarded with build tags.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m53s",
	},
	11: {
		number: 11,
		text:   "Cgo is not Go.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=12m37s",
	},
	12: {
		number: 12,
		text:   "With the unsafe package there are no guarantees.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=13m49s",
	},
	13: {
		number: 13,
		text:   "Clear is better than clever.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s",
	},
	14: {
		number: 14,
		text:   "Reflection is never clear.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=15m22s",
	},
	15: {
		number: 15,
		text:   "Errors are values.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=16m13s",
	},
	16: {
		number: 16,
		text:   "Don't just check errors, handle them gracefully.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=17m25s",
	},
	17: {
		number: 17,
		text:   "Design the architecture, name the components, document the details.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=18m09s",
	},
	18: {
		number: 18,
		text:   "Documentation is for users.",
		url:    "https://www.youtube.com/watch?v=PAAkCSZUG1c&t=19m07s",
	},
	19: {
		number: 19,
		text:   "Don't panic.",
		url:    "https://github.com/golang/go/wiki/CodeReviewComments#dont-panic",
	},
}

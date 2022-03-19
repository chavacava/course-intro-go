# Exercises

## 1 
We refactored our code to use a map instead of a slice to store the proverbs.
How to be sure the refactoring did not intruduce bugs?
Run the tests!
Remember, the command to run the tests is:
```bash
go test 
```

## 2 
Modify the declaration of the `proverbs` map into
```go
var proverbs = map[int]string{
	10: "Cgo must always be guarded with build tags.",
	11: "Cgo is not Go.",
	12: "With the unsafe package there are no guarantees.",
	4:  "The bigger the interface, the weaker the abstraction.",
	5:  "Make the zero value useful.",
	6:  "interface{} says nothing.",
	7:  "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	13: "Clear is better than clever.",
	14: "Reflection is never clear.",
	15: "Errors are values.",
	16: "Don't just check errors, handle them gracefully.",
	17: "Design the architecture, name the components, document the details.",
	18: "Documentation is for users.",
	19: "Don't panic.",
	1:  "Don't communicate by sharing memory, share memory by communicating.",
	2:  "Concurrency is not parallelism.",
	3:  "Channels orchestrate; mutexes serialize.",
	8:  "A little copying is better than a little dependency.",
	9:  "Syscall must always be guarded with build tags.",
}
```
Notice that the order of proverbs has changed.
Run the program and the tests. Everything is OK? Why?

## 3
Modify the declaration of `n` to be
```go
n := 25
```
Run the program. Is the output correct? Why?
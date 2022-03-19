# Exercises

## 1 
Modify the `proverb` structure by adding a new field to store the number of the proverb.

## 2
Add the following instruction at the end of the `main` function
```go
fmt.Printf("\n%v\n", p)
fmt.Printf("\n%v\n", proverbs)
```
Run the program and analyze the output.

## 3
Modify the program to print all the proverbs and their respectives URL.
Expected output:
```
#1      Don't communicate by sharing memory, share memory by communicating. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s)
#2      Concurrency is not parallelism. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=3m42s)
#3      Channels orchestrate; mutexes serialize. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=4m20s)
#4      The bigger the interface, the weaker the abstraction. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s)
#5      Make the zero value useful. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s)
#6      interface{} says nothing. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s)
#7      Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s)
#8      A little copying is better than a little dependency. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s)
#9      Syscall must always be guarded with build tags. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m10s)
#10     Cgo must always be guarded with build tags. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m53s)
#11     Cgo is not Go. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=12m37s)
#12     With the unsafe package there are no guarantees. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=13m49s)
#13     Clear is better than clever. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s)
#14     Reflection is never clear. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=15m22s)
#15     Errors are values. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=16m13s)
#16     Don't just check errors, handle them gracefully. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=17m25s)
#17     Design the architecture, name the components, document the details. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=18m09s)
#18     Documentation is for users. (https://www.youtube.com/watch?v=PAAkCSZUG1c&t=19m07s)
#19     Don't panic. (https://github.com/golang/go/wiki/CodeReviewComments#dont-panic)
```
Tip: use `for-range` intruction
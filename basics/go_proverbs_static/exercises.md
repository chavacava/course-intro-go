# Exercises

## 1 
Modify the code to print the second proverb.
Expected output 
```
There are 19 Go proverbs
Here's one: Concurrency is not parallelism.
```

## 2
Modify the code to print the number of proverbs by modifying the line
```go
fmt.Printf("There are 19 Go proverbs\n")
```
into 
```go
fmt.Printf("There are %d Go proverbs\n", np)
```
Tip: declare a new variable `np` and use the `len` function

## 3
Modify the first line of the function to be
```go
p := proverbs[19]
```
Execute the program.
What is the result? 
Why?

# Exercises

## 1 
Again, our code is buggy, asking for the 10th proverb returns the 11th one. Fix the bug. Expected output 
```
There are 19 Go proverbs
Here's one: #10 Cgo must always be guarded with build tags.
```

## 2
Modify the declaration of `n` to
```go
n := 25
```
Execute the program. What is the output?
Why?

Modify the program to always print a proverb, even in the case `n` is bigger than the actual number of proverbs (19)
You can use the `%` mod operator. 
For example, for `n` being `25` the expexted output is:
```
There are 19 Go proverbs
Here's one: #6 interface{} says nothing.
```

## 3
How we can be sure the function `getProverb` is bug free?

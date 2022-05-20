# Exercises

## 1
The variable `upTo` indicates the number of proverbs to be printed.
Even if it is set to `10` the output contains 11 proverbs.
Why?
Fix the code to print the number of proverbs indicated by `upTo`

## 2
Replace the variable `upTo` by `from`:
```go
from := 10
```
And make the code print from the `from`_th_ proverb up to the last one.
For example, if `from` is `10`, the expected output is:

```
There are 19 Go proverbs
Here are some of them:
#10     Cgo must always be guarded with build tags. 
#11     Cgo is not Go. 
#12     With the unsafe package there are no guarantees. 
#13     Clear is better than clever. 
#14     Reflection is never clear. 
#15     Errors are values. 
#16     Don't just check errors, handle them gracefully. 
#17     Design the architecture, name the components, document the details. 
#18     Documentation is for users. 
#19     Don't panic.
```

## 3
Modify the program to declare both variables `from` and `upTo` and print a range of proverbs.
For example if `from` is `10` and `upTo` is `16` the expected output is:
```
There are 19 Go proverbs
Here are some of them:
#10     Cgo must always be guarded with build tags. 
#11     Cgo is not Go. 
#12     With the unsafe package there are no guarantees. 
#13     Clear is better than clever. 
#14     Reflection is never clear. 
#15     Errors are values. 
#16     Don't just check errors, handle them gracefully. 
```

## 4
Do `from` and `upTo` need to be _variables_? Could we make them _constants_? Try!

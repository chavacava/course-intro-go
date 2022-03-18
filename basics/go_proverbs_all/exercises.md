# Exercises

## 1 
In the output proverbs are numbered from 0 to 18.
Modify the code to number them from 1 to 19.
Expected output 
```
There are 19 Go proverbs:
#1 Don't communicate by sharing memory, share memory by communicating. 
#2 Concurrency is not parallelism. 
#3 Channels orchestrate; mutexes serialize. 
#4 The bigger the interface, the weaker the abstraction. 
#5 Make the zero value useful. 
#6 interface{} says nothing. 
#7 Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. 
#8 A little copying is better than a little dependency. 
#9 Syscall must always be guarded with build tags. 
#10 Cgo must always be guarded with build tags. 
#11 Cgo is not Go. 
#12 With the unsafe package there are no guarantees. 
#13 Clear is better than clever. 
#14 Reflection is never clear. 
#15 Errors are values. 
#16 Don't just check errors, handle them gracefully. 
#17 Design the architecture, name the components, document the details. 
#18 Documentation is for users. 
#19 Don't panic. 
```

## 2 
Proverbs are not aligned.
Modify the printing parameters to print
```
#1      Don't communicate by sharing memory, share memory by communicating. 
#2      Concurrency is not parallelism. 
#3      Channels orchestrate; mutexes serialize. 
#4      The bigger the interface, the weaker the abstraction. 
#5      Make the zero value useful. 
#6      interface{} says nothing. 
#7      Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. 
#8      A little copying is better than a little dependency. 
#9      Syscall must always be guarded with build tags. 
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
Tip: you can use `\t` in the print string

## 3
Modify the code to print the list in reversed order. Expected output:
```
There are 19 Go proverbs:
#19     Don't communicate by sharing memory, share memory by communicating. 
#18     Concurrency is not parallelism. 
#17     Channels orchestrate; mutexes serialize. 
#16     The bigger the interface, the weaker the abstraction. 
#15     Make the zero value useful. 
#14     interface{} says nothing. 
#13     Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. 
#12     A little copying is better than a little dependency. 
#11     Syscall must always be guarded with build tags. 
#10     Cgo must always be guarded with build tags. 
#9      Cgo is not Go. 
#8      With the unsafe package there are no guarantees. 
#7      Clear is better than clever. 
#6      Reflection is never clear. 
#5      Errors are values. 
#4      Don't just check errors, handle them gracefully. 
#3      Design the architecture, name the components, document the details. 
#2      Documentation is for users. 
#1      Don't panic. 
```
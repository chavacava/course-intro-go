# Exercises

## 1 
Run the program as follows
```
go run . -f toto.txt
```
The error is what you expected?

## 2
Run the program as follows
```
go run . -f proverbs_err.txt 
```
The error is what you expected?

## 3
Modify `proverbs_err.txt` to fix the problematic line, replace it with
```
20,Always check slice bounds before accessing it
```  
(let the 2 blank lines at the end of the file)

Run the program as follows
```
go run . -f proverbs_err.txt 
```
What is the error now? Why?
Can you modify the program to ignore lines with no text?

## 4
Modify the program to avoid _global_ (package) variables.

## 5
Modify the program to add two command line flags `from` and `to` and use them to print only those proverbs in the range [`from`-`to`]
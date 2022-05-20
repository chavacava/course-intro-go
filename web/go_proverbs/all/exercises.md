# Exercises

## 1 
Proverbs are printed in random order. Do you know why?
Fix the program to print proverbs in the right order

## 2
Make the output more web-friendly by generating HTML code with web links for each proverb.
For example, the line

> `#2	Concurrency is not parallelism.`

should be replaced by
> `#2`	[`Concurrency is not parallelism`](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=3m42s)`.`

Tips: 
1. the generated output must start with `<html>` and end with `</html>`
1. the HTML code for hyperlinks is 
    ```html
    <a href="url">link text</a>
    ```
1. the HTML tag for line-breaks is `<br/>`

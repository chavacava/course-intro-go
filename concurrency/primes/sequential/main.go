package main

import (
	"fmt"
)

func main() {
	const n = 1000
	primes := primesBetween(1, n)

	fmt.Printf("There are %d primes between 1 and %d:\n%v\n", len(primes), n, primes)
}

func primesBetween(from, to int) []int {
	result := []int{}
	for i := from; i <= to; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

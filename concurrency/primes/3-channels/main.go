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
	candidates := make(chan int)
	// producer
	go func() {
		for i := from; i <= to; i++ {
			candidates <- i
		}
		close(candidates)
	}()

	primes := make(chan int)
	done := make(chan bool)
	const concurrencyLevel = 5
	// checkers
	for i := 1; i <= concurrencyLevel; i++ {
		go func() {
			for candidate := range candidates {
				if isPrime(candidate) {
					primes <- candidate
				}
			}
			done <- true
		}()
	}

	result := []int{}
	doneCheckers := 0
	for {
		select {
		case prime := <-primes:
			result = append(result, prime)
		case <-done:
			doneCheckers++
			if doneCheckers == concurrencyLevel {
				return result
			}
		}
	}
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

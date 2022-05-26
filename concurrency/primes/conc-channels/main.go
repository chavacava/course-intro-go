package main

import (
	"fmt"
)

func main() {
	const n = 1000
	fmt.Printf("The primes between 1 and %d are:\n", n)

	primes := primesBetween(1, n)

	fmt.Printf("%v\n", primes)
}

func primesBetween(from, to int) []int {
	candidates := make(chan int)
	primes := make(chan int)
	done := make(chan bool)

	// producer
	go func() {
		for i := from; i <= to; i++ {
			candidates <- i
		}
		close(candidates)
	}()

	// checkers
	for i := 1; i < 3; i++ {
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
	for {
		select {
		case prime := <-primes:
			result = append(result, prime)
		case <-done:
			return result
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

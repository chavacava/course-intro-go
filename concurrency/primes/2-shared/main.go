package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 1000
	primes := primesBetween(1, n)

	fmt.Printf("There are %d primes between 1 and %d:\n%v\n", len(primes), n, primes)
}

func primesBetween(from, to int) []int {
	result := []int{}
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for i := from; i <= to; i++ {
		wg.Add(1)
		go func(candidate int) {
			if isPrime(candidate) {
				mutex.Lock()
				result = append(result, candidate)
				mutex.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
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

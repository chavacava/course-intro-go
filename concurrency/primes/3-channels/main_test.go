package main

import "testing"

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		n    int
		want bool
	}{
		{
			n:    1,
			want: false,
		},
	}

	for _, tc := range testCases {
		got := isPrime(tc.n)

		if tc.want != got {
			t.Fatalf("expected %d primality to be %v, got %v", tc.n, tc.want, got)
		}
	}
}

func TestPrimesBetween(t *testing.T) {
	testCases := []struct {
		from int
		to   int
		want int // number of expected prime numbers
	}{
		{
			from: 1,
			to:   1,
			want: 0,
		},
		{
			from: 1,
			to:   10,
			want: 4,
		},
		{
			from: 1,
			to:   1000,
			want: 168,
		},
	}

	for _, tc := range testCases {
		got := len(primesBetween(tc.from, tc.to))

		if tc.want != got {
			t.Fatalf("expected %d primes in [%d-%d], got %v", tc.want, tc.from, tc.to, got)
		}
	}
}

func BenchmarkPrimesBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesBetween(1, 100000)
	}
}

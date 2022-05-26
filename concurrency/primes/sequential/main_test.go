package main

import "testing"

func TestIsPrime(t *testing.T) {
	candidate := 1
	got := isPrime(candidate)
	want := false
	if want != got {
		t.Fatalf("expected %d primality to be %v, got %v", candidate, want, got)
	}

	candidate = 2
	got = isPrime(candidate)
	want = true
	if want != got {
		t.Fatalf("expected %d primality to be %v, got %v", candidate, want, got)
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
		primesBetween(1, 1000)
	}
}

package main

import "testing"

func BenchmarkCountChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toUpperCase(proverbs)
	}
}

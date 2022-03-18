package main

import "testing"

func TestGetProverb(t *testing.T) {
	n := 1
	want := "#1 Don't communicate by sharing memory, share memory by communicating."
	got := getProverb(n)

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

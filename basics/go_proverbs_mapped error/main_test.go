package main

import "testing"

func TestGetProverb(t *testing.T) {
	n := 1
	want := "#1 Don't communicate by sharing memory, share memory by communicating."
	wantError := ""
	got, err := getProverb(n)

	if wantError != "" {
		// we expect an error, lets check there is an error
		if err == nil {
			t.Fatalf("Expected error %q, got nothing", wantError)
		}
		// we got an error, lets check if it is the one we want
		if err.Error() != wantError {
			t.Fatalf("Expected error %q, got: %q", wantError, err.Error())
		}
	} else if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}

	// if we add more test cases this method will be too long...
}

package main

import "testing"

func TestGetProverb(t *testing.T) {
	type testCase struct {
		n         int
		want      string
		wantError string
	}

	cases := []testCase{
		{
			n:         1,
			want:      "#1 Don't communicate by sharing memory, share memory by communicating.",
			wantError: "",
		},
		{
			n:         19,
			want:      "#19 Don't panic.",
			wantError: "",
		},
		// add more test cases here
	}

	for _, tc := range cases {
		got, err := getProverb(tc.n)
		if tc.wantError != "" {
			// we expect an error, lets check there is an error
			if err == nil {
				t.Fatalf("Expected error %q, got nothing", tc.wantError)
			}
			// we got an error, lets check if it is the one we want
			if err.Error() != tc.wantError {
				t.Fatalf("Expected error %q, got: %q", tc.wantError, err.Error())
			}
		} else if got != tc.want {
			t.Fatalf("Expected %q, got %q", tc.want, got)
		}
	}

}

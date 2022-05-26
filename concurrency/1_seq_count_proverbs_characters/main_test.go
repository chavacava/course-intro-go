package main

import "testing"

func TestCountChars(t *testing.T) {
	testCases := map[string]struct {
		sentences []string
		want      int
	}{
		"empty": {
			sentences: []string{},
			want:      0,
		},
	}

	for name, tc := range testCases {
		got := countChars(tc.sentences)

		if tc.want != got {
			t.Fatalf("test case %q: expected %d, got %d", name, tc.want, got)
		}
	}
}

func BenchmarkCountChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countChars(proverbs)
	}
}

// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {

	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{name: "empty", input: "", want: 0},
		{name: "a2", input: "a2_32_^_&", want: 1},
		{name: "812", input: "812_%6ab", want: 2},
	}

	lc := letterCounter{}

	// range over the tests and run them as subtests
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := lc.count(tc.input)

			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count

// write a test for symbolCounter.count

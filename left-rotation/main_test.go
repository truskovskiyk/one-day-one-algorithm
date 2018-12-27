package main

import "testing"

func TestRotLeft(t *testing.T) {

	n := 8
	path := "UDDDUDUU"
	answer := 1
	result := countingValleys(n, path)
	if result != answer {
		t.Errorf("rotLeft was incorrect, got: %v, want: %v.", result, answer)
	}
}



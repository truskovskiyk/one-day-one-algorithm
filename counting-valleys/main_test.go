package main

import "testing"

func TestCountingValleys(t *testing.T) {
	n := 8
	path := "UDDDUDUU"
	answer := 1
	result := countingValleys(n, path)
	if result != answer {
		t.Errorf("countingValleys was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestCountingValleys2(t *testing.T) {
	n := 8
	path := "DDUUUUDD"
	answer := 1
	result := countingValleys(n, path)
	if result != answer {
		t.Errorf("countingValleys was incorrect, got: %v, want: %v.", result, answer)
	}
}
package main

import "testing"

func testSmoke(t *testing.T) {
	result := 1
	answer := 1
	if result != answer {
		t.Errorf("gameOfThrones was incorrect, got: %v, want: %v.", result, answer)
	}
}
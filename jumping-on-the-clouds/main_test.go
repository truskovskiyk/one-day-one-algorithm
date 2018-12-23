package main

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	if result != answer {
		t.Errorf("jumpingOnClouds was incorrect, got: %v, want: %v.", result, answer)
	}
}
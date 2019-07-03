package main

import "testing"

func TestTwoStringsCase1(t *testing.T) {

	result := twoStrings("hello", "world")
	answer := "YES"
	if result != answer {
		t.Errorf("twoStrings was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestTwoStringsCase2(t *testing.T) {

	result := twoStrings("hi", "world")
	answer := "NO"
	if result != answer {
		t.Errorf("twoStrings was incorrect, got: %v, want: %v.", result, answer)
	}
}

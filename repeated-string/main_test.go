package main

import "testing"

func TestRepeatedString(t *testing.T) {
	s := "aba"
	n := int64(10)

	answer := int64(7)
	result := repeatedString(s, n)
	if result != answer {
		t.Errorf("repeatedString was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestRepeatedStringHuge(t *testing.T) {
	s := "a"
	n := int64(1000000000000)

	answer := int64(1000000000000)
	result := repeatedString(s, n)
	if result != answer {
		t.Errorf("repeatedString was incorrect, got: %v, want: %v.", result, answer)
	}
}

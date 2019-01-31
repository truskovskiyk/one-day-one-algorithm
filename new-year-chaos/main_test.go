package main

import "testing"


func TestMinimumBribesResultTractable(t *testing.T) {
	q := []int{2, 1, 5, 3, 4}
	result := minimumBribesResult(q)
	answer := "3"
	if result != answer {
		t.Errorf("minimumBribesResult was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestMinimumBribesResultInTractable(t *testing.T) {
	q := []int{2, 5, 1, 3, 4}
	result := minimumBribesResult(q)
	answer := "Too chaotic"
	if result != answer {
		t.Errorf("minimumBribesResult was incorrect, got: %v, want: %v.", result, answer)
	}
}

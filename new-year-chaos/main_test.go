package main

import (
	"testing"
	"fmt"
)


func TestMinimumBribesResultTractableCase1(t *testing.T) {
	q := []int{2, 1, 5, 3, 4}
	result := minimumBribesResult(q)
	fmt.Print(result)
	answer := "3"
	if result != answer {
		t.Errorf("minimumBribesResult was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestMinimumBribesResultInTractableCase1(t *testing.T) {
	q := []int{2, 5, 1, 3, 4}
	result := minimumBribesResult(q)
	answer := "Too chaotic"
	if result != answer {
		t.Errorf("minimumBribesResult was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestMinimumBribesResultInTractableCase2(t *testing.T) {
	q := []int{5, 1, 2, 3, 7, 8, 6, 4}
	result := minimumBribesResult(q)
	answer := "Too chaotic"
	if result != answer {
		t.Errorf("minimumBribesResult was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestMinimumBribesResultTractableCase2(t *testing.T) {
	q := []int{1, 2, 5, 3, 7, 8, 6, 4}
	result := minimumBribesResult(q)
	fmt.Print(result)
	answer := "7"
	if result != answer {
		t.Errorf("minimumBribesResult was incorrect, got: %v, want: %v.", result, answer)
	}
}
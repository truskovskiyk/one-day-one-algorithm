package main

import "testing"

func isIntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRotLeft(t *testing.T) {

	d := 4
	arr := []int{1, 2, 3, 4, 5}
	answer := []int{5, 1, 2, 3, 4}
	result := rotLeft(arr, d)
	if !isIntArrayEquals(result, answer) {
		t.Errorf("rotLeft was incorrect, got: %v, want: %v.", result, answer)
	}
}

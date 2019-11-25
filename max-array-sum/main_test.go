package main

import "testing"


func TestMaxSubsetSum(t *testing.T) {

	arr := []int{3, 7, 4, 6, 5}
	answer := 11
	result := maxSubsetSum(arr)
	if result != answer {
		t.Errorf("MaxSubsetSum was incorrect, got: %v, want: %v.", result, answer)
	}
}

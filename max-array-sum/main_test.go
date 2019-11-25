package main

import "testing"


func TestMaxSubsetSum1(t *testing.T) {

	arr := []int64{-2, 1, 3, -4, 5}
	answer := int64(8)
	result := maxSubsetSum(arr)
	if result != answer {
		t.Errorf("MaxSubsetSum was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestMaxSubsetSum2(t *testing.T) {

	arr := []int64{3, 7, 4, 6, 5}
	answer := int64(13)
	result := maxSubsetSum(arr)
	if result != answer {
		t.Errorf("MaxSubsetSum was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestMaxSubsetSum3(t *testing.T) {

	arr := []int64{2, 1, 5, 8, 4}
	answer := int64(11)
	result := maxSubsetSum(arr)
	if result != answer {
		t.Errorf("MaxSubsetSum was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestMaxSubsetSum4(t *testing.T) {

	arr := []int64{3, 5, -7, 8, 10}
	answer := int64(15)
	result := maxSubsetSum(arr)
	if result != answer {
		t.Errorf("MaxSubsetSum was incorrect, got: %v, want: %v.", result, answer)
	}
}
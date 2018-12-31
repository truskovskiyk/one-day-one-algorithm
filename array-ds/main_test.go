package main

import "testing"

func TestHourglassSum(t *testing.T) {
	arr := [][]int{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}
	result := hourglassSum(arr)

	answer := 19
	if result != answer {
		t.Errorf("hourglassSum was incorrect, got: %v, want: %v.", result, answer)
	}
}

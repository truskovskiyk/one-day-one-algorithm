package main

import (
	"testing"
	"fmt"
)



func TestMinimumSwapsCase1(t *testing.T) {
	arr := []int{2, 1, 5, 3, 4}


	result := MinimumSwaps(arr)
	fmt.Print(result)
	answer := int32(3)
	if ! (result == answer) {
		t.Errorf("MinimumSwaps was incorrect, got: %v, want: %v.", result, answer)
	}
}



func TestMinimumSwapsCase2(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4, 6, 7}


	result := MinimumSwaps(arr)
	fmt.Print(result)
	answer := int32(3)
	if ! (result == answer) {
		t.Errorf("MinimumSwaps was incorrect, got: %v, want: %v.", result, answer)
	}
}
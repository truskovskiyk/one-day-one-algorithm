package main

import "testing"

func TestSockMerchant(t *testing.T) {
	n := 9
	arr := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	answer := 3
	result := sockMerchant(n, arr)
	if result != answer {
		t.Errorf("sockMerchant was incorrect, got: %v, want: %v.", result, answer)
	}
}

package main

import "testing"

func testEq(a, b []int32) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCutTheSticks(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 3, 3, 2, 1}
	asnwer := []int32{8, 6, 4, 1}
	res := CutTheSticks(arr)
	if !testEq(res, asnwer) {
		t.Errorf("CutTheSticks was incorrect, got: %v, want: %v.", res, asnwer)
	}
}

package main

import "testing"

func TestArrayManipulationCase1(t *testing.T) {
	queries := [][]int32{

		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},

	}
	n := int32(5)
	result := arrayManipulation(n, queries)
	answer := int64(200)
	if result != answer {
		t.Errorf("arrayManipulation was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestArrayManipulationCase2(t *testing.T) {
	queries := [][]int32{

		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},

	}
	n := int32(10)
	result := arrayManipulation(n, queries)
	answer := int64(10)
	if result != answer {
		t.Errorf("arrayManipulation was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestArrayManipulationCase3(t *testing.T) {
	queries := [][]int32{

		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{5, 9, 15},


	}
	n := int32(10)
	result := arrayManipulation(n, queries)
	answer := int64(31)
	if result != answer {
		t.Errorf("arrayManipulation was incorrect, got: %v, want: %v.", result, answer)
	}
}

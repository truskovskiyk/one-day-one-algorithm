package main

import "testing"

func TestGameOfThronesCase2(t *testing.T) {
	n := 4
	m := 4
	k := 3

	track := [][]int{{2, 2, 3}, {3, 1, 4}, {4, 4, 4}}
	result := GridlandMetro(n, m, k, track)
	answer := 9
	if result != answer {
		t.Errorf("TestGameOfThronesCase1 was incorrect, got: %v, want: %v.", result, answer)
	}

}

func TestGridlandMetroCase1(t *testing.T) {
	n := 2
	m := 8
	k := 1

	track := [][]int{
		{2, 2, 3},
		{1, 3, 6},
		{1, 2, 7},
	}
	result := GridlandMetro(n, m, k, track)
	answer := 8
	if result != answer {
		t.Errorf("TestGridlandMetroCase1 was incorrect, got: %v, want: %v.", result, answer)
	}

}

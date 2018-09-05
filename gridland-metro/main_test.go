package main

import "testing"

func TestGameOfThronesCase1(t *testing.T) {
	n := int32(4)
	m := int32(4)
	k := int32(3)

	track := [][]int32{}
	result := GridlandMetro(n, m, k, track)
	answer := int32(9)
	if result != answer {
		t.Errorf("TestGameOfThronesCase1 was incorrect, got: %v, want: %v.", result, answer)
	}

}

package main

import "testing"

//func TestGameOfThronesCase1(t *testing.T) {
//	n := int32(4)
//	m := int32(4)
//	k := int32(3)
//
//	track := [][]int32{{2, 2, 3}, {3, 1, 4}, {4, 4, 4}}
//	result := GridlandMetro(n, m, k, track)
//	answer := int32(9)
//	if result != answer {
//		t.Errorf("TestGameOfThronesCase1 was incorrect, got: %v, want: %v.", result, answer)
//	}
//
//}


func TestGameOfThronesCase2(t *testing.T) {
	n := int32(1)
	m := int32(8)
	k := int32(1)

	track := [][]int32{
		{1, 2, 3},
		{1, 3, 6},
		{1, 2, 7},

		}
	result := GridlandMetro(n, m, k, track)
	answer := int32(1)
	if result != answer {
		t.Errorf("TestGameOfThronesCase1 was incorrect, got: %v, want: %v.", result, answer)
	}

}
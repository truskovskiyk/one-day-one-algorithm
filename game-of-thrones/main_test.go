package main

import "testing"

func testSimpleInput(s string, answer string, t *testing.T) {
	result := gameOfThrones(s)
	if result != answer {
		t.Errorf("gameOfThrones was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestGameOfThronesCase1(t *testing.T) {
	testSimpleInput("aaabbbb", "YES", t)
}

func TestGameOfThronesCase2(t *testing.T) {
	testSimpleInput("cdefghmnopqrstuvw", "NO", t)
}

func TestGameOfThronesCase3(t *testing.T) {
	testSimpleInput("cdcdcdcdeeeef", "YES", t)
}

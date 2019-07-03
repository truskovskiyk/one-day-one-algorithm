package main

import "testing"

func TestCheckMagazineCase1(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}


	result := checkMagazine(magazine, note)
	answer := "Yes"
	if result != answer {
		t.Errorf("checkMagazine was incorrect, got: %v, want: %v.", result, answer)
	}
}


func TestCheckMagazineCase2(t *testing.T) {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}


	result := checkMagazine(magazine, note)
	answer := "No"
	if result != answer {
		t.Errorf("checkMagazine was incorrect, got: %v, want: %v.", result, answer)
	}
}

package main

import "testing"

func TestJumpingOnCloudsOneCloud(t *testing.T) {
	answer := 0
	clouds := []int{0}
	result := jumpingOnClouds(clouds)
	if result != answer {
		t.Errorf("jumpingOnClouds was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestJumpingOnCloudsTwoClouds(t *testing.T) {
	answer := 1
	clouds := []int{0, 0}
	result := jumpingOnClouds(clouds)
	if result != answer {
		t.Errorf("jumpingOnClouds was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestJumpingOnCloudsCase1(t *testing.T) {
	answer := 4
	clouds := []int{0, 0, 1, 0, 0, 1, 0}
	result := jumpingOnClouds(clouds)
	if result != answer {
		t.Errorf("jumpingOnClouds was incorrect, got: %v, want: %v.", result, answer)
	}
}

func TestJumpingOnCloudsCase2(t *testing.T) {
	answer := 3
	clouds := []int{0, 0, 0, 1, 0, 0}
	result := jumpingOnClouds(clouds)
	if result != answer {
		t.Errorf("jumpingOnClouds was incorrect, got: %v, want: %v.", result, answer)
	}
}

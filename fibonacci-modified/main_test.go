package main

import "testing"

func TestFibonacciModified(t *testing.T) {
	total := FibonacciModified(10, 0, 1)
	if total != "84266613096281243382112" {
		t.Errorf("FibonacciModified was incorrect, got: %v, want: %v.", total, "84266613096281243382112")
	}
}

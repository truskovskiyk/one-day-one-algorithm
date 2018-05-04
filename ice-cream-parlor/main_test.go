package main

import "testing"

func TestBSNotFound(t *testing.T) {
	total := BinarySearch(10, []int{1, 2, 3, 4, 5, 6})
	if total != -1 {
		t.Errorf("BinarySearch was incorrect, got: %v, want: %v.", total, -1)
	}
}

func TestBSFound(t *testing.T) {
	total := BinarySearch(3, []int{1, 2, 3, 4, 5, 6})
	if total != 2 {
		t.Errorf("BinarySearch was incorrect, got: %v, want: %v.", total, 2)
	}
}

func TestBSFoundNotRange(t *testing.T) {
	total := BinarySearch(3, []int{1, 3, 432, 3455})
	if total != 1 {
		t.Errorf("BinarySearch was incorrect, got: %v, want: %v.", total, 1)
	}
}

func TestIcecreamParlor(t *testing.T) {
	m := 4
	c := []int{1, 4, 5, 3, 2}
	firstIndex, secondIndex := IcecreamParlor(m, c)
	if firstIndex != 0 || secondIndex != 3{
		t.Errorf("IcecreamParlor was incorrect , got: (%v, %v), want: (%v, %v).", firstIndex, secondIndex, 0, 3)
	}
}

func TestIcecreamParlorSameElements(t *testing.T) {
	m := 4
	c := []int{2, 2, 4, 3}
	firstIndex, secondIndex := IcecreamParlor(m, c)
	if firstIndex != 0 || secondIndex != 1{
		t.Errorf("IcecreamParlor was incorrect , got: (%v, %v), want: (%v, %v).", firstIndex, secondIndex, 0, 1)
	}
}
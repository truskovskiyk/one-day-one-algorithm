package main

import (
	"testing"
)

func Eq(a, b []int32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
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

func TestBFS(t *testing.T) {
	n := int32(4)
	m := int32(2)
	edges := [][]int32{{1, 2}, {1, 3}}
	s := int32(1)
	result := bfs(n, m, edges, s)
	resultExpected := []int32{6, 6, -1}
	if !Eq(result, resultExpected) {
		t.Errorf("bfs was incorrect, got: %v, want: %v.", result, resultExpected)
	}

	n = int32(3)
	m = int32(1)
	edges = [][]int32{{2, 3}}
	s = int32(2)
	result = bfs(n, m, edges, s)
	resultExpected = []int32{-1, 6}
	if !Eq(result, resultExpected) {
		t.Errorf("bfs was incorrect, got: %v, want: %v.", result, resultExpected)
	}
}

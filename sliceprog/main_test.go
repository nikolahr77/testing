package main

import "testing"

func TestGetTheMinElement(t *testing.T) {
	min := GetMinElement([]int{4, 3, 6, -7, 2, 8})
	if min != -7 {
		t.Errorf("Min element was wrong, got %d, want %d", min, -7)
	}
}

func TestGetTheSumOfSliceElements(t *testing.T) {
	sum := GetSum([]int{4, 3, 6, -7, 2, 8})
	if sum != 16 {
		t.Errorf("Sum was incorrect, got %d, want %d", sum, 16)
	}
}

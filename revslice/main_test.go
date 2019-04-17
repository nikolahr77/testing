package main

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	s := ReverseSlice([]int{4, 3, 6, -7, 2, 8})
	expected := []int{8, 2, -7, 6, 3, 4}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("The slice was not reversed correctly, want %d, got %d", expected, s)
	}
}

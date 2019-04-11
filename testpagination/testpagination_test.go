package testpagination

import (
	"reflect"
	"testing"
)

func TestPages_FirstPageAndSetCurrentToFirst(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat", "house"}, PageSize: 3}
	actual := s.FirstPage()
	expected := []string{"book", "work", "cat"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FirstPage was incorrect, got: %d, want %d", actual, expected)
	}
	if s.current != 1 {
		t.Errorf("Current page was not set, got: %d, want %d", s.current, 1)

	}
}

func TestPages_LastPageAndSetCurrentToLast(t *testing.T) {
	s := Pages{Elements: []string{"house", "phone", "case", "cup", "clock", "dog"}, PageSize:3}
	last := s.LastPage()
	result := []string{ "cup","clock","dog"}

	if !reflect.DeepEqual(last, result) {
		t.Errorf("LastPage was incorrect, got: %d, want %d.", last, result)
	}
	if s.current != 2 {
		t.Errorf("Current page was not set, got: %d, want %d", s.current, 2)
	}
}
func TestPages_LastPageAndSetCurrentToLastWithLessElements(t *testing.T) {
	s := Pages{Elements: []string{"house", "phone", "case", "cup", "clock", "dog","horn","book"}, PageSize:3}
	last := s.LastPage()
	result := []string{ "horn","book"}

	if !reflect.DeepEqual(last, result) {
		t.Errorf("LastPage was incorrect, got: %d, want %d.", last, result)
	}
	if s.current != 3 {
		t.Errorf("Current page was not set, got: %d, want %d", s.current, 3)
	}
}

func TestPages_GetCurrentPageNumber(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat"}, PageSize: 1,current:2}

	if s.GetCurrentPageNumber() != 2 {
		t.Errorf("Current page was not returned, got: %d, want %d", s.GetCurrentPageNumber(), 3)
	}
}

func TestPages_NextPage(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat", "house"}, PageSize: 3}

	if !s.HasNext() {
		t.Errorf("There should be a next page,got %t, want %t.", s.HasNext(), true)
	}
}

func TestPages_NextPageWhenThereIsNoNextPage(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "flat", "house","junk"}, PageSize: 2,current:3}

	if s.HasNext() {
		t.Errorf("There should not be a next page,got %t, want %t.", s.HasNext(), true)
	}
}

func TestPages_PreviousPageExists(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat", "house","cup"}, PageSize: 2,current:2}

	if !s.HasPrevious()  {
		t.Errorf("There should be a previous page,got %t, want %t.", s.HasPrevious(), true)
	}
}

func TestPages_PreviousPageDoesntExist(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat", "house"}, PageSize: 2}

	if s.HasPrevious()  {
		t.Errorf("There should not be a previous page,got %t, want %t.", s.HasPrevious(), true)
	}
}

func TestPages_GetNextPage(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat", "house"}, PageSize: 2,current:1}
	next := s.Next()
	result := []string{"cat", "house"}

	if !reflect.DeepEqual(next, result) {
		t.Errorf("Next didn't return the next page, got %d, want %d", next, result)
	}
}

func TestPages_GetNextPageDoesntExist(t *testing.T) {
	s := Pages{Elements: []string{"book", "work", "cat", "house","pen"}, PageSize: 2,current:3}
	next := s.Next()
	if next != nil {
		t.Error("There should not be a next page to return, next it not nil")
	}
}

func TestPages_GetPreviousPage(t *testing.T) {
	s := Pages{Elements: []string{"book", "phone", "cat", "house","pen"}, PageSize: 2,current:3}
	previous := s.Previous()
	result := []string{"cat", "house"}

	if reflect.DeepEqual(previous, result) != true {
		t.Errorf("Previous didn't return the previous page, got %d, want %d", previous, result)
	}
}

func TestPages_GetPreviousPageDoesntExist(t *testing.T) {
	s := Pages{Elements: []string{"book", "phone", "cat", "house","pen"}, PageSize: 2,current:1}
	previous := s.Previous()

	if previous != nil {
		t.Error("There should not be a previous page to return, previous is not nil")
	}
}

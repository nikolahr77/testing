package main

import "fmt"

//ReverseSlice returns a slice in reverse order
func ReverseSlice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		help := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = help
	}
	return s
}

func main() {
	s := []int{4, 3, 6, -7, 2, 8}

	fmt.Print(ReverseSlice(s))
}

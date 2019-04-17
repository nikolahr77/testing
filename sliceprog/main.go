package main

import "fmt"

// GetMinElement returns the lowest number in a slice
func GetMinElement(s []int) int {
	help := s[0]
	for _, v := range s {
		if help > v {
			help = v
		}
	}
	return help
}

// GetSum returns all the elements in a slice
func GetSum(s []int) int {
	sum := 0
	for _, v := range s {
		sum = sum + v
	}
	return sum
}

func main() {
	slice := []int{4, 3, 6, -7, 2, 8}

	fmt.Println(GetMinElement(slice))
	fmt.Println(GetSum(slice))
	fmt.Println(slice)
}

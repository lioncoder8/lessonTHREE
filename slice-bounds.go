package main

import "fmt"

func main_array_slice() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

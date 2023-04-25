package main

import "fmt"

func mainSlce_len_go() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Println("len = %d, cap = %d %v\n", len(s), cap(s), s)
}

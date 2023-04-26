package main

import "fmt"

func main_nil() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

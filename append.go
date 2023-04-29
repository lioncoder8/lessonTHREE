package main

import "fmt"

func main_append() {
	var s []int
	printSlicea(s)

	s = append(s, 0)
	printSlicea(s)
}

func printSlicea(x []int) {
	fmt.Printf("lend = %d cap = %d %v\n", len(x))
}

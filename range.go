package main

import "fmt"

var pow = []int{1, 2, 3, 4, 16, 32, 64, 128}

func main_range() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

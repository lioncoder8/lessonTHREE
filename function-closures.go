package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return i
	}
}

func main_function_closures() {
	pos, neg := adder(), adder()
	for x := 0; x < 10; x++ {
		fmt.Println(
			pos(x),
			neg(-2*x),
		)
	}
}

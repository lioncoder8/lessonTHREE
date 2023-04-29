package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev1 := 0
	prev2 := 1
	return func() int {
		fib := prev1 + prev2
		prev1 = prev2
		prev2 = fib
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

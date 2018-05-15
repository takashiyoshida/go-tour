package main

import "fmt"

// fibonacci is a function that returns a function that returns an int
func fibonacci() func() int {
	// 0, 1, 1, 2, 3, 5, 8, ...
	a, b, c := 0, 1, 0
	return func() int {
		// we can compute the next fibonacci value but will return the current one
		c, a, b = a, b, a+b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

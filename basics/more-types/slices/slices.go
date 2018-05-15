package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// returns the 1st element (remember it is zero indexed)
	// but excludes the 4th one; [3, 5, 7]
	var s []int = primes[1:4]
	fmt.Println(s)
}

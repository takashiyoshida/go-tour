package main

import "fmt"

func Sqrt(x float64) float64 {
	// starting with a guess z, adjust z based on how close z^2 is to x,
	// producing a better guess
	z := 1.0 // initial guess is 1
	// repeat the calculation 10 times
	// print each z

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("guess (%d): %g\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

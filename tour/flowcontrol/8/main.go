// Exercise: Loops and Functions.
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	prev := z
	var i int
	for i = 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if z == prev || math.Abs(z-prev) < 0.0000001 {
			break
		}
		prev = z

		fmt.Printf("\t%f\n", z)
	}
	fmt.Printf("\tCount: %d\n", i)
	return z
}

func main() {
	for x := float64(0); x < 10; x++ {
		fmt.Println("x", x)
		fmt.Printf("\tResult:    %f\n", Sqrt(x))
		fmt.Printf("\tmath.Sqrt: %f\n", math.Sqrt(x))
	}
}

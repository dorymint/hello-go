// Exercise: Errors.
package main

import (
	"fmt"
	"math"
)

// Expected: ErrNegativeSqrt(-2).Error() then
// "cannot Sqrt negative number: -2"
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

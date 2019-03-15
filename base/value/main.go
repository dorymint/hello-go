// huge number.
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := uint(0); i < 65; i++ {
		fmt.Printf("%3d: %d %d %d %d\n", i, int(0), ^int(0), int(1)<<i, ^(int(1) << i))
	}
	fmt.Println("math.MaxInt64:", math.MaxInt64)
}

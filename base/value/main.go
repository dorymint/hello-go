package main

// TODO: impl

import (
	"fmt"
	"math"
)

func main() {
	for i := uint(0); i < 70; i++ {
		fmt.Printf("%3d: %d %d %d %d\n", i, int(0), ^int(0), int(1)<<i, ^(int(1)<<i))
	}
	fmt.Println("max int64:", math.MaxInt64)
}

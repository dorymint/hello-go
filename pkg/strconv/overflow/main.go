// pkg/strconv/overflow.
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Printf("maxint: %d\n", math.MaxInt64)
	maxint := strconv.FormatInt(math.MaxInt64, 10)

	v, err := strconv.ParseInt(maxint+"0", 10, 64)
	fmt.Printf("v:%v\n", v)
	fmt.Printf("err:%v\n", err)
}

// bin.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	const b = 1<<63 - 1
	fmt.Println(b)
	bigInt := new(big.Int).SetInt64(b)
	fmt.Println("bigInt:", new(big.Int).Add(bigInt, bigInt))

	x := new(big.Int)
	y := new(big.Int)
	x.SetInt64(10)
	y.SetInt64(30)
	fmt.Println("x,y", x, y)

	z := new(big.Int).Add(x, y)
	fmt.Println("z", z)
	z.Add(z, z)
	fmt.Println("z", z)

	fmt.Println("Exp x x", new(big.Int).Exp(x, x, nil))
}

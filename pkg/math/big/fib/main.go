package main

import (
	"fmt"
	"math/big"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func helloBig() {
	const b = 1 << 63 -1
	fmt.Println(b)
	bigInt := new(big.Int).SetInt64(b)
	fmt.Println("bigInt:", new(big.Int).Add(bigInt,bigInt))

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

func fib() {
	fib := make([]*big.Int, 100)
	for i := range fib {
		fib[i] = new(big.Int)
	}
	fib[0].SetInt64(0)
	fib[1].SetInt64(1)
	for i := 2; i < len(fib); i++ {
		fib[i].Add(fib[i-2], fib[i-1])
	}
	for i := 0; i < len(fib); i++ {
		fmt.Println(fib[i])
	}
}

func main() {
	split("helloBig")
	helloBig()
	split("fib")
	fib()
}

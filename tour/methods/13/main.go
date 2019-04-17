// Nil interface values.
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)

	defer func() { fmt.Println(recover()) }()
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

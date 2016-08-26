package main

import (
	"fmt"
)

func pointer_1() {
	fmt.Println("pointer_1")
	i, j := 42,	2701

	p := &i
	fmt.Println("&p = ", &p)
	fmt.Println("p  = ", p)
	fmt.Println("&i = ", &i)
	fmt.Println(*p)

	fmt.Println()

	p = &j
	fmt.Println("&p = ", &p)
	fmt.Println("p  = ", p)
	fmt.Println("&j = ", &j)
	fmt.Println(*p)
	return
}

func pointer_2() {
	fmt.Println("pointer_2")
	i, j := 42,	2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	fmt.Println()

	p = &j
	*p = *p / 37
	fmt.Println(j)

	return
}

func main() {

	pointer_1()
	fmt.Println()
	pointer_2()

	// NOTE:ポインタ演算はない

}

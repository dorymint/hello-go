package main

import (
	"fmt"
)

func main() {
	fmt.Println("nil length")
	{
		var n []string
		fmt.Println("len",
			len(n),
			len([]string{}),
			len(*new([]string)),
		)
	}

	fmt.Printf("\nslice caps\n")
	{
		n := []string{"hello", "world"}
		fmt.Println("len",
			len(n),
			len(n[1:]),
			"cap",
			cap(n),
			cap(n[1:]),
		)
	}

	fmt.Printf("\nslice pointer\n")
	{
		n := []int{1, 2, 3}
		fmt.Printf("pointer n =%p &n =%p\n", n, &n)
		n2 := n[1:]
		fmt.Printf("pointer n2=%p &n2=%p\n", n2, &n2)
		fmt.Println("is same?", &n == &n2)

		fmt.Printf("&n[1] %p\n", &n[1])
		fmt.Printf("n2    %p\n", n2)
	}
}

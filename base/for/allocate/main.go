// base/for/allocate.
package main

import "fmt"

func NewErr(i int) error {
	return fmt.Errorf("%d", i)
}

func main() {
	fmt.Println(":=")
	for i := 0; i < 10; i++ {
		if err := NewErr(i); err != nil {
			fmt.Printf("%v %p %p\n", err, &err, err)
		}
	}

	fmt.Println("=")
	var err error
	for i := 0; i < 10; i++ {
		if err = NewErr(i); err != nil {
			fmt.Printf("%v %p %p\n", err, &err, err)
		}
	}
}

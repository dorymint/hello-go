// base/defer/variable.
package main

import "fmt"

func modAtReturn() *int {
	p := new(int)
	var i int
	*p = i
	defer func() { *p = 1 }()
	return p
}

func main() {
	fmt.Printf("modAtReturn: %d\n", *modAtReturn())
}

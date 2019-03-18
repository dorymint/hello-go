// boolean.
package main

import "fmt"

func main() {
	var b bool = false
	fmt.Printf("%v\n", b)

	b = !b
	fmt.Printf("%v\n", b)

	b = !b
	fmt.Printf("%v\n", b)
}

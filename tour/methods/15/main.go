// Type assertions.
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	defer func() { fmt.Println(recover()) }()
	f = i.(float64) // panic
	fmt.Println(f)
}

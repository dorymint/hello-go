// Exercise: Fibonacci closure.
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i int
	var next int = 1
	return func() int {
		defer func() {
			i, next = next, i+next
		}()
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

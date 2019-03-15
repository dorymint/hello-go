// scope.
package main

import "fmt"

func main() {
	i := func(i int) { fmt.Println(i) }
	i(1)
	{
		defer i(2) // call at return
		i(3)
	}
	i(4)

	return

	// Output:
	// 1
	// 3
	// 4
	// 2
}

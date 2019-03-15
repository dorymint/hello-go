// for.
package main

import (
	"fmt"
)

func main() {
	for a, n := 'a', 'n'; a != 'n'; {
		fmt.Println(string(a), string(n))
		a++
		n++
	}
}

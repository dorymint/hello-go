// iterate.
package main

import (
	"fmt"
)

func main() {
	ss := []string{"hello", "gopher", "world"}
	for i, s := range ss {
		fmt.Println(i, s)
	}
}

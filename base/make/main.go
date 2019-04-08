// make.
package main

import "fmt"

func main() {
	n := -1
	defer func() {
		fmt.Println(recover())
	}()
	_ = make([]int, 0, n)
}

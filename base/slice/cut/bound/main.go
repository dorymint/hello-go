// bound of length.
package main

import "fmt"

func main() {
	ss := []string{}
	defer func() {
		fmt.Println(recover())
	}()
	_ = ss[1:]
}

// uint overflow.
package main

import "fmt"

func main() {
	ui := uint(0)
	fmt.Println(ui - 100) // overflow, may print huge number
}

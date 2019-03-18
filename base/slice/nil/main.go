// nil.
package main

import "fmt"

func main() {
	var ss []string
	fmt.Printf("ss==nil:%v\n", ss == nil)
	fmt.Println("len", len(ss))
	fmt.Println(ss[:] == nil)
}

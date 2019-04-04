// usage of +build.
//
//	go run ./
//
package main

import "fmt"

var Name string = "not specified"

func main() {
	fmt.Printf("os %s\n", Name)
}

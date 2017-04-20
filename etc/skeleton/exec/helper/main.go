package main

import (
	"fmt"
)

// From stackoverflow
// http://stackoverflow.com/questions/23166468/how-can-i-get-stdin-to-exec-cmd-in-golang

func main() {
	fmt.Println("Hello, What's your favorite number?")
	var i int
	fmt.Scanf("%d\n", &i)
	fmt.Println("Ah I like ", i, "too.")
}

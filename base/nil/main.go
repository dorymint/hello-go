// nil.
package main

import "fmt"

func main() {
	var p *string

	// not panic
	if p == nil || *p == "s" {
		fmt.Println("not panic:")
	}

	// panic
	defer func() {
		fmt.Printf("%+v\n", recover())
	}()
	if *p == "s" || p == nil {
		fmt.Println("panic: can not display")
	}
}

// parse.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var v int
	flag.IntVar(&v, "v", 0, "parse example")
	fmt.Println(v)

	// input
	os.Args = []string{"cmdname", "-v", "100"}
	flag.Parse()
	fmt.Println(v)

	os.Args = []string{"cmdname", "-v", "want error"}

	// unreachable
	defer fmt.Println(recover())

	// call os.Exit in flag.Parse
	flag.Parse()

	// unreachable
	fmt.Println(v)
}

// readdirnames.

// +build linux

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(".")
	if err != nil {
		panic(err)
	}

	names, err := f.Readdirnames(-1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", names)
}

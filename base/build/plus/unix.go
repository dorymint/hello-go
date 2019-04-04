// only unix.

// +build unix

package main

import "fmt"

func init() {
	Name = "unix"
	fmt.Println("from unix")
}

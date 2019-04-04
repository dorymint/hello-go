// only linux.

// +build linux

package main

import "fmt"

func init() {
	Name = "linux"
	fmt.Println("from linux")
}

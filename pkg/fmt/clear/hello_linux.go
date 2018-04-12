// +build linux

package main

import (
	"fmt"
)

// Hi from linux
func Hi() {
	fmt.Println("hello from: hello_linux.go")
}

// HiDup
func HiDup() {
	fmt.Println("HiDup from: hello_linux.go")
}


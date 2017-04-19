// +build windows

package main

import (
	"fmt"
)

// Hi from windows
func Hi() {
	fmt.Println("Hi from: hello_windows.go")
}

// HiDup
func HiDup() {
	fmt.Println("HiDup from: hello_windows.go")
}

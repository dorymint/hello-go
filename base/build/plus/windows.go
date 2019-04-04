// only windows.

// +build windows

package main

import "fmt"

func init() {
	Name = "windows"
	fmt.Println("only windows")
}

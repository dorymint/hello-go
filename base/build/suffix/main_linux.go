// *_linux.go is build only linux.
// on linux, if exists *_unix.go then merge to sources.
package main

import "fmt"

func init() {
	Name = "linux"
	fmt.Println("from linux")
}

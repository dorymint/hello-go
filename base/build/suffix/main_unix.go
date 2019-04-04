// *_unix.go is build only unix.
package main

import "fmt"

func init() {
	Name = "unix"
	fmt.Println("from unix")
}

package main

import (
	"fmt"
	"path"
)

func clean() {
	fmt.Println("clean:")

	unix := `/../test`
	fmt.Println("unix:", unix, "to", path.Clean(unix))

	win := `c:\windows\..\path`
	fmt.Println("windows:", win, "to", path.Clean(win))
	// out: windows: c:\windows\..\path to c:\windows\..\path
}

func main() {
	ss := []string{
		"hello",
		"world",
	}
	fmt.Println(path.Join(ss...))
	ss2 := []string{
		"http://hello///world/",
		"testbase.html",
	}
	fmt.Println(path.Join(ss2...))
	// trimed // >> /
	// http:/hello/world/testbase.html

	clean()
}

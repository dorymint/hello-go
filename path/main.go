package main

import (
	"fmt"
	"path"
)

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
}

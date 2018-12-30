package main

// TODO: impl

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")
	ss := []string{
		"/",
		"/view/hello",
		"/edit/hello",
		"/save/hello",
		"/view/hello/world",
		"/v/",
		"../fef/",
		string(0x00),
	}

	repf := func(src []byte) []byte { return []byte("hello") }

	for i, s := range ss {
		out := reg.ReplaceAllFunc([]byte(s), repf)
		fmt.Println(i, s, string(out))
	}
}

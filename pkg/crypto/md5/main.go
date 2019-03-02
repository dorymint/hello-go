package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
)

func main() {
	var h hash.Hash
	h = md5.New()
	fmt.Printf("nul:\t%x\n", h.Sum(nil))

	fmt.Printf("hello1:\t%x\n", h.Sum([]byte("hello")))
	io.WriteString(h, "hello")
	fmt.Printf("hello2:\t%x\n", h.Sum(nil))

	fmt.Printf("world1:\t%x\n", h.Sum([]byte("world")))
	io.WriteString(h, "world")
	fmt.Printf("world2:\t%x\n", h.Sum(nil))

	h.Reset()
	fmt.Printf("Reset():\t%x\n", h.Sum(nil))
}

package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	sha := sha512.New512_256()

	fmt.Println("sha.Sum(nil):", sha.Sum(nil))
	fmt.Printf("%x\n", sha.Sum(nil))

	sha.Write([]byte("hello world"))

	fmt.Println("sha.Sum(nil):", sha.Sum(nil))
	fmt.Printf("%x\n", sha.Sum(nil))
}

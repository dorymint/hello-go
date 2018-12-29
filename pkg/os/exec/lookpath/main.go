package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("sh")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	fmt.Printf("%T\n", path)
}

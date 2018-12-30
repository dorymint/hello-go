package main

// TODO: impl

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("PID:", os.Getpid())
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Name(), w.Name())
}

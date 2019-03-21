// open.
//
//	go run main.go
//
package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := io.Copy(os.Stdout, f); err != nil {
		panic(err)
	}
}

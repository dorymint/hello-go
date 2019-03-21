package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	file := filepath.Join("testdata", "go.txt")
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(b))
}

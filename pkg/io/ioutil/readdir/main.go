package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dir := "testdata"

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, fi := range fis {
		fmt.Printf("%#v\n", fi)
		fmt.Printf("fi.Name():%q\n\n", fi.Name())
	}
}

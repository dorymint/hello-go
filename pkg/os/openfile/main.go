// open file.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	tempf, err := ioutil.TempFile("", "openfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempf.Name())

	err = ioutil.WriteFile(tempf.Name(), []byte("test openfile"), 0600)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(tempf.Name(), os.O_RDWR|os.O_TRUNC, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("rewritten")

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("size %v\n", fi.Size())
	fmt.Printf("parm %v\n", fi.Mode().Perm())
}

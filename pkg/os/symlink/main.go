// symlink.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	testroot, err := ioutil.TempDir("", "test_symlink")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(testroot)

	dir, err := ioutil.TempDir(testroot, "test_symlink")
	if err != nil {
		panic(err)
	}

	sym := dir + ".link"
	err = os.Symlink(dir, sym)
	if err != nil {
		panic(err)
	}

	fis, err := ioutil.ReadDir(testroot)
	if err != nil {
		panic(err)
	}
	for i, fi := range fis {
		fmt.Printf("%d\n", i)
		fmt.Printf("\t%s\n", fi.Name())
		fmt.Printf("\t%+v\n", fi.Mode())
	}
}

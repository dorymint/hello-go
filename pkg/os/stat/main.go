// stat.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var (
		file    = "file"
		symlink = "symlink"
	)

	stat := func(file string) {
		fi, err := os.Stat(file)
		if err != nil {
			panic(err)
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name:%q\n", fi.Name())
		fmt.Printf("Mode:%v\n", fi.Mode())
		fmt.Printf("Parm:%v\n", fi.Mode().Perm())
		fmt.Printf("Content:%q\n", content)
		fmt.Printf("IsSymlink:%v\n\n", fi.Mode()&os.ModeSymlink != 0)
	}

	stat(file)
	stat(symlink)
}

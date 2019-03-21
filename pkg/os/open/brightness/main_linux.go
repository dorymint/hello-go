// cat brightness.

// +build linux

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	const (
		dir     = "/sys/class/backlight"
		current = "brightness"
		max     = "max_brightness"
	)

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var rootes []string
	for _, fi := range fis {
		if fi.Mode()&os.ModeSymlink != 0 {
			fi, err = os.Stat(filepath.Join(dir, fi.Name()))
			if err != nil {
				panic(err)
			}
		}
		if fi.IsDir() {
			rootes = append(rootes, filepath.Join(dir, fi.Name()))
		}
	}

	cat := func(root, file string) {
		f, err := os.Open(filepath.Join(root, current))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := io.Copy(os.Stdout, f); err != nil {
			panic(err)
		}
	}

	for _, root := range rootes {
		fmt.Printf("root:%#v\n", root)
		fmt.Println("current brightness")
		cat(root, current)

		fmt.Println("max brightness")
		cat(root, max)
		fmt.Println()
	}
}

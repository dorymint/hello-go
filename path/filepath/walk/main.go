package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var root = flag.String("root", "", "")

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("invalid args:", flag.Args())
	}
	var err error
	*root, err = filepath.Abs(*root)
	if err != nil {
		log.Fatal(err)
	}

	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("\n\n%v\n\n", err)
		}
		fmt.Printf("path:%s:", path)
		if info != nil {
			fmt.Printf("%p:", info)
		}
		if err != nil {
			fmt.Printf("err:%v", err)
		}
		fmt.Println()
		return nil
	}
	err = filepath.Walk(*root, f)
	if err != nil {
		log.Fatal(err)
	}
}

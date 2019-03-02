package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	flag.Parse()
	fmt.Println(flag.CommandLine.Name())
	fmt.Println(filepath.Base(flag.CommandLine.Name()))

	fs := flag.NewFlagSet("new flag set", flag.ExitOnError)
	fs.Parse(nil)
	fmt.Println(fs.Name())
}

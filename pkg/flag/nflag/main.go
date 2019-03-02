package main

import (
	"flag"
	"fmt"
)

func main() {
	help := flag.Bool("help", false, "Display help")
	version := flag.Bool("version", false, "Display version")
	exit := flag.Bool("exit", false, "Exit")

	flag.Parse()

	fmt.Println("narg", flag.NArg())
	fmt.Println("nflag", flag.NFlag())

	switch {
	case *help:
		fmt.Println("help")
	case *version:
		fmt.Println("version")
	case *exit:
		return
	}
}

package main

import (
	"fmt"
	"os"
)

// self parse

func ckargs(args []string, wantNArg int) {
	if len(args) == wantNArg {
		return
	}
	fmt.Fprintf(os.Stderr, "unexpected arguments: %v\n", args)
	os.Exit(1)
}

func main() {
	fmt.Printf("args: %q\n", os.Args)

	var args []string
	var arg string

	for args = os.Args[1:]; len(args) != 0; args = args[1:] {
		arg = args[0]

		switch arg {
		case "-v", "-version", "--version":
			ckargs(args, 1)
			fmt.Println("version")
			return
		case "-h", "-help", "--help":
			ckargs(args, 1)
			fmt.Println("help")
			return
		default:
		}
	}

	fmt.Printf("exit: len(args)=%d\n", len(args))
}

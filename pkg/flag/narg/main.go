// NAlag.
package main

import (
	"flag"
	"fmt"
)

func main() {
	fs := flag.NewFlagSet("nargtest", flag.ExitOnError)
	_ = fs.Bool("version", false, "print version")

	args := []string{"-version", "hello", "world"}
	if err := fs.Parse(args); err != nil {
		panic(err)
	}

	fmt.Printf("NFlag:%d\n", fs.NArg())
}

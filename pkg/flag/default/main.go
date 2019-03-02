package main

import (
	"flag"
	"fmt"
)

func main() {
	var s string
	flag.StringVar(&s, "string", "default", "")
	flag.StringVar(&s, "s", "default", "Alias of -string")
	flag.Parse()
	flag.Usage()

	fmt.Printf("s=%s\n", s)
}

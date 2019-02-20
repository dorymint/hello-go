package main

import (
	"fmt"
)

func main() {
	// Ref:
	// $ godoc -http=localhost:6060
	// $ $BROWSER localhost:6060/pkg/fmt
	const (
		indent = "\t"
		nl     = "\n"
		// ...
	)
	fmt.Printf("hello\thello world\n")
	fmt.Printf("hello%shello world%s", indent, nl)
}

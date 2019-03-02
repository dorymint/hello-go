package main

import (
	"flag"
	"fmt"
)

var flagvar int

func main() {
	flag.IntVar(&flagvar, "flagvar", 0, "error test flagvar")
	fmt.Println(flagvar)

	// Input:
	// -flagvar=errortest
	flag.Parse() // Parse error, runtimeerror
	// os.Exit
	fmt.Println(flagvar) // this do not run
}

// make usage.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const Name = "cmdname"

func makeUsage() (*io.Writer, func()) {
	var w io.Writer = os.Stderr
	usage := func() {
		flag.CommandLine.SetOutput(w)
		// two spaces
		fmt.Fprintf(w, "Description:\n")
		fmt.Fprintf(w, "  Short description\n\n")
		fmt.Fprintf(w, "Usage:\n")
		fmt.Fprintf(w, "  %s [Options]\n\n", Name)
		fmt.Fprintf(w, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(w, "Examples:\n")
		examples := `
  $ ` + Name + ` -help # Display help message
`
		fmt.Fprintf(w, "%s\n", examples)
	}
	return &w, usage
}

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "print the version number")
	flag.Parse()

	fmt.Fprintln(os.Stderr, "--- Default usage ---")
	flag.Usage() // to stderr

	var w *io.Writer
	w, flag.Usage = makeUsage()
	*w = os.Stdout
	fmt.Println("--- Custom usage ---")
	flag.Usage() // to stdout

	if version {
		fmt.Printf("%s %s\n", Name, "0.0.0")
	}
}

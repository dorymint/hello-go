package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var flagColor = flag.Bool("color", false, "")

func init() {
	flag.Parse()
	color.NoColor = *flagColor
}

func main() {
	color.Cyan("hello")
	color.Blue("hello")
	color.Red("hello")

	fmt.Fprintln(os.Stdout, color.MagentaString("hello string"))
}

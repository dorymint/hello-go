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

	fmt.Println("color list")
	fns := []func(string, ...interface{}){
		func(s string, i ...interface{}) { fmt.Println("----- Colors -----") },
		color.Black,
		color.Red,
		color.Green,
		color.Yellow,
		color.Blue,
		color.Magenta,
		color.Cyan,
		color.White,

		func(s string, i ...interface{}) { fmt.Println("----- Hi Colors -----") },
		color.HiBlack,
		color.HiRed,
		color.HiGreen,
		color.HiYellow,
		color.HiBlue,
		color.HiMagenta,
		color.HiCyan,
		color.HiWhite,
	}

	for _, f := range fns {
		f("colors !!")
	}
}

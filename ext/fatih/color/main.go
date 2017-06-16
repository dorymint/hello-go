package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var flagColor = flag.Bool("color-off", false, "")

func init() {
	flag.Parse()
	color.NoColor = *flagColor
}

func main() {
	color.Cyan("hello")
	color.Blue("hello")
	color.Red("hello")

	fmt.Fprintln(os.Stdout, color.MagentaString("hello string"))

	fmt.Println("----- color list -----")
	fns := []struct {
		f func(string, ...interface{})
		c string
	}{
		{func(s string, i ...interface{}) { fmt.Println("----- Colors -----") }, ""},
		{color.Black, "Black"},
		{color.Red, "Red"},
		{color.Green, "Green"},
		{color.Yellow, "Yellow"},
		{color.Blue, "Blue"},
		{color.Magenta, "Magenta"},
		{color.Cyan, "Cyan"},
		{color.White, "White"},

		{func(s string, i ...interface{}) { fmt.Println("----- Hi Colors -----") }, ""},
		{color.HiBlack, "HiBlack"},
		{color.HiRed, "HiRed"},
		{color.HiGreen, "HiGreen"},
		{color.HiYellow, "HiYellow"},
		{color.HiBlue, "HiBlue"},
		{color.HiMagenta, "HiMagenta"},
		{color.HiCyan, "HiCyan"},
		{color.HiWhite, "HiWhite"},
	}

	for _, f := range fns {
		f.f(f.c)
	}
}

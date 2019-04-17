// ANSI escape code.
package main

import (
	"fmt"
)

// \x1b はESCの16進表記

type color string

const reset color = "\x1b[0m"

// foreground 文字色 \x1b[30m ... 37m ...default=39m
const (
	fgBlack   color = "\x1b[30m"
	fgRed     color = "\x1b[31m"
	fgGreen   color = "\x1b[32m"
	fgYellow  color = "\x1b[33m"
	fgBlue    color = "\x1b[34m"
	fgMagenta color = "\x1b[35m"
	fgCyan    color = "\x1b[36m"
	fgWhite   color = "\x1b[37m"

	fgDefault color = "\x1b[39m"
)

// background 背景色 \x1b[40m ... 47m ...default=49m
const (
	bgBlack   color = "\x1b[40m"
	bgRed     color = "\x1b[41m"
	bgGreen   color = "\x1b[42m"
	bgYellow  color = "\x1b[43m"
	bgBlue    color = "\x1b[44m"
	bgMagenta color = "\x1b[45m"
	bgCyan    color = "\x1b[46m"
	bgWhite   color = "\x1b[47m"

	bgDefault color = "\x1b[49m"
)

// 装飾
const (
	emphasis  color = "\x1b[1m"
	underline color = "\x1b[4m"
	reverse   color = "\x1b[7m"
)

var fgList = []color{
	fgBlack,
	fgRed,
	fgGreen,
	fgYellow,
	fgBlue,
	fgMagenta,
	fgCyan,
	fgWhite,
	fgDefault,
}

var bgList = []color{
	bgBlack,
	bgRed,
	bgGreen,
	bgYellow,
	bgBlue,
	bgMagenta,
	bgCyan,
	bgWhite,
	bgDefault,
}

var decorateList = []color{
	emphasis,
	underline,
	reverse,
}

// ターミナルにカラーコードを付けて送る
func wrap(s string, c color) {
	fmt.Print(string(c) + s + string(reset))
}

func main() {
	printColors := func(cs []color) {
		for i := range cs {
			wrap("hello world", cs[i])
			fmt.Println()
		}
	}

	fmt.Println("fgColors")
	printColors(fgList)
	fmt.Println()

	fmt.Println("bgColors")
	printColors(bgList)
	fmt.Println()

	fmt.Println("decorates")
	printColors(decorateList)
	fmt.Println()
}

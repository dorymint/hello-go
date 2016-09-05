
package main

// terminal color

import (
	"fmt"
)

func colorTest() {
	// \x1b はESCの16進表記

	// foreground 文字色 \x1b[30m ... 37m ...default=39m
	const (
		fgBlack		string = "\x1b[30m"
		fgRed		string = "\x1b[31m"
		fgGreen		string = "\x1b[32m"
		fgYellow	string = "\x1b[33m"
		fgBlue		string = "\x1b[34m"
		fgMagenta	string = "\x1b[35m"
		fgCyan		string = "\x1b[36m"
		fgWhite		string = "\x1b[37m"
		fgDefault	string = "\x1b[39m"
	)

	// background 背景色 \x1b[40m ... 47m ...default=49m
	const (
		bgBlack		string = "\x1b[40m"
		bgRed		string = "\x1b[41m"
		bgGreen		string = "\x1b[42m"
		bgYellow	string = "\x1b[43m"
		bgBlue		string = "\x1b[44m"
		bgMagenta	string = "\x1b[45m"
		bgCyan		string = "\x1b[46m"
		bgGray		string = "\x1b[47m"
		bgDefault	string = "\x1b[49m"
	)

	// 装飾
	const (
		colorDefault	string = "\x1b[0m"
		emphasis		string = "\x1b[1m"
		underline		string = "\x1b[4m"
		reverse			string = "\x1b[7m"
	)


	fgList := []string{
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

	bgList := []string{
		bgBlack,
		bgRed,
		bgGreen,
		bgYellow,
		bgBlue,
		bgMagenta,
		bgCyan,
		bgGray,
		bgDefault,
	}

	decorateList := []string{
		colorDefault,
		emphasis,
		underline,
		reverse,
	}

	// ターミナルにカラーコードを送る
	sendColor := func(s string) { fmt.Printf("%s", s) }
	// view
	view := func(s []string){
		for i := 0; i < len(s); i++ {
			sendColor(s[i])
			fmt.Printf("hello world")
			sendColor(colorDefault)
			fmt.Printf("\n")
		}
	}

	/* Test */

	fmt.Printf("fgColor Test\n")
	view(fgList)

	fmt.Printf("\n\nbgColor Test\n")
	view(bgList)

	fmt.Printf("\n\nbgColor Test2\n")
	sendColor(fgWhite)
	for i := 0; i < len(bgList); i++ {
		// TODO:表示の挙動が怪しい
		// ターミナルの画面をリセットしてgo run
		// hello worldの部分だけ背景色が付く
		// そのままもう一度go run
		// 改行部分も含めて色が付く
		sendColor(bgList[i])
		fmt.Println("hello world")
		sendColor(bgDefault)
		fmt.Printf("\n")
	}
	sendColor(colorDefault)

	fmt.Printf("\n\ndecorate Test\n")
	view(decorateList)

	sendColor(colorDefault)
	fmt.Println("end of color test")

	/* end of Test */

}

func main() {

	colorTest()

}

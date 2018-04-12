package main

import (
	"flag"
	"fmt"
	"os"
)

func printRed(str string) {
	fmt.Print("\x1b[31m")
	fmt.Print(str)
	fmt.Print("\x1b[0m")
}
func printBlue(str string) {
	fmt.Print("\x1b[34m")
	fmt.Print(str)
	fmt.Print("\x1b[0m")
}

func helloFlag() {
	printRed("helloFlag\n")
	f := flag.Int("flag1", 0, "this is flag1 help msg")
	fmt.Println("flag.Args:", flag.Args())
	flag.Parse() // flag1にint以外を引数に渡してるとここparseを読んだ時にランタイムエラーになる
	//flag.Parse() // 二回目以降は?

	// パースを読んだ後フラグではない引数が残る
	fmt.Println("flag Args")
	fmt.Println(flag.Args())
	fmt.Println(flag.Arg(0), flag.Arg(1), flag.Arg(2))

	fmt.Println(os.Args)
	for i, x := range os.Args {
		fmt.Println("os.Args", i, ":", x)
	}

	fmt.Println("flag.Args:", flag.Args())
	if *f == 100 {
		fmt.Println("hello!")
	}

	fmt.Println("flag1")
	flag.PrintDefaults()
}

var flagvar int

func init() {
	printRed("init\n")
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	// 1,変数へのバインド 2,デフォルト 3,ヘルプ

	// flag.Parse()を呼んでいないので反映されない
	printBlue(fmt.Sprintln("flagvvar:", flagvar))
}

func flagUsage() {
	printRed("flagUsage\n")
	var ip = flag.Int("flagname1", 1234, "help message for flagname")
	// var ip2 = flag.Int("flagname1", 1234, "help message for flagname")
	// 同じフラグネームを投げるとランタイムパニック

	fmt.Println("ip:", ip, *ip)
	printBlue(fmt.Sprintln("flagvar:", flagvar))

	fmt.Println("flag usage")
	flag.PrintDefaults()
	// show help
}

func parseTest() {
	flag.Parse()
	fmt.Println(flagvar) // flagvar=入力された値
	flagvar = 0
	fmt.Println(flagvar) // 0
	flag.Parse()         // 二回目以降もパースされる
	fmt.Println(flagvar) // flagvar=入力された値

	// NOTE:
	// flag.Parse()はフラグ引数の値も検証する
	// -flagvar=errortest
	// としているとParseが呼ばれた時に初めてランタイムエラーになる
}

func main() {
	printRed("main\n")

	helloFlag()
	fmt.Println("main flag usage")
	flagUsage()

	fmt.Println("main")
	flag.PrintDefaults()

	parseTest()
}

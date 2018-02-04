package main

import (
	"flag"
	"fmt"
	"strings"
)

func q1() {
	var f = flag.String("f", "flag", "")
	fmt.Println(*f)
	flag.Parse()
	fmt.Println(*f)
	// run: go run main.go hello
	// out:flag\nhello
}

func q2() {
	var mainflag = flag.String("mf", "hello", "")
	var main2flag = flag.String("mf2", "world", "")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*mainflag, *main2flag)

	fmt.Println("NArg:", flag.NArg())
	// can't parsed arguments
	fmt.Println("NFlag:", flag.NFlag())
	// number of flag arguments
}

// go run main.go
// go run main.go -mainf=world
// go run main.go sub -subf=hello
// go run main.go -mainf=world sub -subf=hello
// go run main.go -h
// go run main.go sub -h
func q3() {
	var mainflag = flag.String("mainf", "hello", "")
	var sub *flag.FlagSet
	var subflag *string
	// main flag parse
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println(*mainflag)
		return
	}
	// sub flag parse
	switch flag.Arg(0) {
	case "echo":
		fmt.Println(flag.Args())
	case "sub":
		sub = flag.NewFlagSet(strings.Join(flag.Args(), " "), flag.ExitOnError)
		subflag = sub.String("subf", "world", "")
		err := sub.Parse(flag.Args()[1:])
		if err != nil {
			fmt.Println("sub.Parse:", err)
		}
	default:
		fmt.Println("mainflag: ", *mainflag)
		return
	}
	str := "subflag: " + *subflag + "\n" +
		"sub.NFlag: " + string(sub.NFlag()) + "\n" +
		"sub.NArg: " + string(sub.NArg()) + "\n" +
		"sub.Arg: " + strings.Join(sub.Args(), " ")
	fmt.Println(str)
	fmt.Println("mainf: ", *mainflag)
}

func main() {
	q3()
}

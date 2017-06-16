package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func q1() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

L:
	for {
		fmt.Print("tty:>")
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		switch r {
		case 'q':
			break L
		default:
			fmt.Println(string(r))
		}
	}
}

func q2() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	fmt.Print("read path:>")
	s, err := tty.ReadPassword()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("input:", s)
}

func q3() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	fmt.Print("read path clear:>")
	s, err := tty.ReadPasswordClear()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Println("input:", s)
}

func main() {
	split("q1")
	q1()

	split("q2")
	q2()

	split("q3")
	q3()
}

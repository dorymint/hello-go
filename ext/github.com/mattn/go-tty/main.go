package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

func readRune() {
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

func readPassword() {
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

func readPasswordClear() {
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
	fmt.Println("input:", s)
}

func main() {
	readRune()
	readPassword()
	readPasswordClear()
}

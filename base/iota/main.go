// iota.
package main

import (
	"fmt"
	"os"
)

type ExitCode int

const (
	Valid ExitCode = iota
	ErrInit
	ErrFatal
)

func (ec ExitCode) Int() int {
	return int(ec)
}

func run() ExitCode {
	sendValid := make(chan ExitCode)
	go func() {
		sendValid <- Valid
	}()
	sendErrInit := make(chan ExitCode)
	go func() {
		sendErrInit <- ErrInit
	}()
	sendErrFatal := make(chan ExitCode)
	go func() {
		sendErrFatal <- ErrFatal
	}()

	var exit ExitCode
	select {
	case exit = <-sendValid:
	case exit = <-sendErrInit:
	case exit = <-sendErrFatal:
	}
	return exit
}

func main() {
	fmt.Println("random exit code")
	exit := run()
	fmt.Println(exit)
	os.Exit(exit.Int())
}

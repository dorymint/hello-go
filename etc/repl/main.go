// etc/repl.
// repl(Read-eval-print loop)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLine() (string, error) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if err := sc.Err(); err != nil {
		return "", err
	}
	return sc.Text(), nil
}

func printHelp() {
	fmt.Println(`Commands:
	help     Display this message
	exit, :q Exit`)
}

func run() error {
	printHelp()
loop:
	for {
		fmt.Print("Input $")
		text, err := readLine()
		if err != nil {
			return err
		}
		text = strings.TrimSpace(text)
		switch text {
		case "exit", ":q":
			break loop
		case "help":
			printHelp()
		default:
			fmt.Println(text)
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

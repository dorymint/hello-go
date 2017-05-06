package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// impl repl
// make subcommands

// SubCommands interp functions for Repl
type SubCommands map[string]func() (string, error)

// ErrValidExit for valid exit, use Repl, return nil
var ErrValidExit = errors.New("valid exit")

func exit() (string, error) {
	return "", ErrValidExit
}

// Repl Read Eval P.*? Loop
// call function in SubCommands[string]
// string is from os.Stdin
// if return ErrValidExit then return nil
func (sub SubCommands) Repl(prefix string) error {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prefix)
		if !sc.Scan() {
			return fmt.Errorf("fail sc.Scan")
		}
		if sc.Err() != nil {
			return sc.Err()
		}
		f, ok := sub[strings.TrimSpace(sc.Text())]
		if !ok {
			fmt.Printf("invalid subcommand: %q\n", sc.Text())
			continue
		}
		result, err := f()
		if err != nil {
			switch err {
			case ErrValidExit:
				return nil
			default:
				return err
			}
		}
		fmt.Println(result)
	}
}

// SubNew return SubCommands with included interpreter for exit
// "exit" "quit" ":q"
// them are call the valid exit, return nil
func SubNew() SubCommands {
	sub := make(SubCommands)
	sub["exit"] = exit
	sub["quit"] = exit
	sub[":q"] = exit
	return sub
}

// Gomem test gomem
type Gomem struct {
	title   string
	content string
}

func (g *Gomem) echo() (string, error) {
	fmt.Println("echo.tilte=", g.title)
	fmt.Println("echo.content=", g.content)
	g.title = "after"
	return "title changed", nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	sub := SubNew()
	g := &Gomem{title: "before", content: "content"}
	sub["echo"] = g.echo
	if err := sub.Repl("repl:>"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("g=", g) // content after change

	gn := new(Gomem)
	fmt.Println("Before: gn=", gn)
	s, err := gn.echo()
	fmt.Println("After:  gn=", s, err)
}

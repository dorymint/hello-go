package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// ErrExit for call os exit
var ErrExit = errors.New("exit")

// ErrUndifineSubCmd for continue
var ErrUndifineSubCmd = errors.New("unknown subcommand")

// Command for interactive
type Command struct {
	Name string
	Args []string
}


// TODO: implement
func(c Command) Run() error {
	switch c.Name {
	case "exit":
		return ErrExit
	case "echo":
		if len(c.Args) < 1 {
			return errors.New("echo: invalid args")
		}
		echo(c.Args[0])
	default:
		return ErrUndifineSubCmd
	}
	return nil
}

// TODO: add commands
// consider: interactive
func makeCommand(cmdName string, args []string) Command {
	var c Command
	c.Args = args
	switch cmdName {
	case "exit":
		c.Name = "exit"
	case "echo":
		c.Name = "echo"
		if c.Args == nil {
			c.Args = append(c.Args, read("echo:>"))
		}
	default:
		c.Name = cmdName
	}
	return c
}

func read(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		return ""
	}
	return sc.Text()
}

func interactive() error {
	fmt.Print("repl:>")
	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			return sc.Err()
		}

		s := strings.TrimSpace(sc.Text())
		c := makeCommand(s, nil)
		err := c.Run()
		if err != nil {
			switch err {
			case ErrExit:
				return nil
			case ErrUndifineSubCmd:
				fmt.Println(sc.Text())
			default:
				return err
			}
		}
		fmt.Print("repl:>")
	}
	return errors.New("interactive: undefined error, fail sc.Scan()")
}

func main() {
	log.SetFlags(log.Lshortfile)
	if err := interactive(); err != nil {
		log.Fatal(err)
	}
	c := makeCommand("echo", []string{"hi", "hi", "hi", })
	c.Run()
}

// test mock echo
func echo(s string) {
	fmt.Println(s)
}

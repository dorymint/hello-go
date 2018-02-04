package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var i = flag.Bool("interactive", false, "")

type command struct {
	name string
	args []string
}

func main() {
	flag.Parse()
	if *i {
		if err := repl(); err != nil {
			log.Fatal(err)
		}
		return
	}

	c := argsParse(flag.Args())
	switch c.name {
	case "exit":
		fmt.Fprintln(os.Stderr, "exit")
		return
	default:
		fmt.Fprintln(os.Stderr, "invalid subcommand:", c)
		// return
	}
	fmt.Println(c)

	fmt.Println("hi")
	fmt.Printf("%q\n", flag.Args())
}


/// mocks

func repl() error {
	return nil
}

func argsParse(args []string) (command) {
	if args == nil {
		return command{name: "exit"}
	}
	c := command{
		args: make([]string, len(args)),
	}
	for i, v := range args {
		if i == 0 {
			c.name = v
			continue
		}
		c.args = append(c.args, v)
	}
	return c
}

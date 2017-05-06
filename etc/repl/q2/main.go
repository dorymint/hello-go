package main

import (
	"errors"
)

// Gomem type memo
type Gomem struct {
	title   string
	content string
	tags    []string
}

// Gomems list
type Gomems map[string]Gomem

// Command is Definition sub commands
type Command struct {
	Name string
	Args []string

	f     func() error
	gomem *Gomems
}

// Commands is map to command
type Commands map[string]Command

// Run from c.Name
func (c Command) Run() error {
	if c.f == nil {
		return errors.New("c.Run invalid: c.f == nil")
	}
	switch c.Name {
	case "echo":
		//do echo(c.Args)
		println("impl?: do echo(c.Args)")
	default:
		return errors.New("undefined command name")
	}
	return errors.New("c.Run invalid: undefined function")
}

var commands = make(Commands)

func init() {
	// definition commands
	commands["echo"] = Command{
		Name: "echo",
		Args: []string{},
	}
}

func main() {
	if err := commands["echo"].Run(); err != nil {
		panic(err)
	}
}

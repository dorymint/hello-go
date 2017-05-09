package subcmd

/*
Example:
	sub := New()
	// puts function pointer
	sub["string"] = func() (string, error)
	// run interactive: read eval print loop
	err := sub.Repl("specify prefix")
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// SubCommands interp functions for Repl
type SubCommands map[string]func() (string, error)

// ErrValidExit for valid exit, for Repl
var ErrValidExit = errors.New("valid exit")

// Repl is Read Eval Print Loop
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

// Example: for valid exit on the Repl
func exit() (string, error) {
	return "", ErrValidExit
}

// NewWithExit return SubCommands with included interpreter for exit
// "exit" "quit" ":q"
// them call the valid exit, return nil
func NewWithExit() SubCommands {
	sub := make(SubCommands)
	sub["exit"] = exit
	sub["quit"] = exit
	sub[":q"] = exit
	return sub
}

// +build linux

package main

import (
	"os"
	"os/exec"
)

// clear display
func clear() error {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	return c.Run()
}

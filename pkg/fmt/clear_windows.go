// +build windows

package main

import (
	"os"
	"os/exec"
)


func clear() error {
	c := exec.Command("cls")
	c.Stdout = os.Stdout
	return c.Run()
}

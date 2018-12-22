package main

import (
	"fmt"
	"os/exec"
)

func main() {
	ss := []string{"pwd"}
	cmd := exec.Command(ss[0], ss[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

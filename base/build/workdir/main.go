// git.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var s = func() string {
	b, err := exec.Command("git", "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		if workdir, err := os.Getwd(); err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "work directory:%s\n", workdir)
		}
		panic(err)
	}
	return strings.TrimSpace(string(b))
}()

func main() {
	fmt.Println(s)
}

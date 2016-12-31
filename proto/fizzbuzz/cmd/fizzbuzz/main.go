package main

import (
	fb	"../../"
	"fmt"
	"os"
)


func main() {
	if err := fb.FizzBuzz(150); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

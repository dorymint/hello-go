package main

import (
	fb	"github.com/dory/hello-go/fizzbuzz"
	"fmt"
	"os"
)


func main() {
	if err := fb.FizzBuzz(150); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

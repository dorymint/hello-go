package main

import (
	"./fizzbuzz"

	"fmt"
	"os"
)


func main() {
	if err := fizzbuzz.FizzBuzz(150); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

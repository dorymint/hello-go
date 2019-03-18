// rune.
package main

import "fmt"

func main() {
	var line string = "name arg1 arg2"
	for i := range line {
		r := rune(line[i])
		fmt.Printf("%+4v %4s\n", r, string(r))
	}
}

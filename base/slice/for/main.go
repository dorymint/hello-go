// for.
package main

import "fmt"

func main() {
	line := "name arg1 arg2"
	for i := range line {
		fmt.Printf("%d %q %q\n", i, string(line[i]), line[i+1:])
	}

	// not panic
	fmt.Printf("\nline[len(line):]\n")
	fmt.Printf("not panic, %%q=%q\n", line[len(line):])
}

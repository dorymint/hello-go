// Type inference.
package main

import "fmt"

func main() {
	print := func(name string, v interface{}) {
		fmt.Printf("%s is of type %T\n", name, v)
	}

	v := true
	print("v", v)

	i := 42
	print("i", i)

	f := 3.142
	print("f", f)

	g := 0.867 + 0.5i
	print("g", g)
}

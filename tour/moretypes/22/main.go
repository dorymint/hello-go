// Mutating Maps.
package main

import "fmt"

func main() {
	m := make(map[string]int)
	const key = "Answer"

	answer := func() {
		v, ok := m[key]
		fmt.Println("The value:", v)
		if !ok {
			fmt.Println("The value:", v, "Present?", ok)
		}
	}

	m[key] = 42
	answer()

	m[key] = 48
	answer()

	delete(m, key)
	answer()
}

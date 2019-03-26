// append.
package main

import "fmt"

func main() {
	var st = struct{ ss []string }{}

	fmt.Printf("%#v\n", st)
	st.ss = append(st.ss, "hello")
	fmt.Printf("%#v\n", st)
}

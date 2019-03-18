// Fields.
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	exp := []string{"name", "arg1", "arg2"}
	fmt.Printf("expected arguments: %q\n\n", exp)

	ck := func(line string) {
		fmt.Printf("in: %q\n", line)
		out := strings.Fields(line)
		fmt.Printf("out: %q\n", out)
		fmt.Printf("exp==out: %v\n\n", reflect.DeepEqual(exp, out))
	}

	ck("name arg1 arg2")
	ck(" name arg1	arg2		")
	ck(`cmd "arg1 arg2" arg3`)
	ck(`cmd."arg1.arg2".arg3`)
}

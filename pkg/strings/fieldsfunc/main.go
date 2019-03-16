// FieldsFunc.
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	exp := []string{"name", "arg1", "arg2"}
	ck := func(line string) {
		out := strings.FieldsFunc(line, func(r rune) bool { return r == '.' })
		fmt.Printf("in: %q\n", line)
		fmt.Printf("out: %q\n", out)
		fmt.Printf("out is expected: %v\n\n", reflect.DeepEqual(exp, out))
	}

	ck("name.arg1.arg2")
	ck(" name.arg1.arg2 ")
	ck("name.\"arg1.arg2\".arg2")
}

// pkg/reflect/typeof.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var integer int
	var i interface{}
	i = integer
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	default:
		panic("unreachable")
	}

	fmt.Println("reflect.TypeOf:", reflect.TypeOf(i))
}

// atoi.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "128"
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v: %d\n", reflect.TypeOf(i), i*2)
}

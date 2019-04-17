// pkg/reflect/deepequal.
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string `json:"name"`
}

func main() {
	t := &T{"foo"}
	var nt interface{} = &T{"foo"}

	fmt.Println(reflect.DeepEqual(t, nt))
}

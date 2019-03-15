// override.
package main

import (
	"encoding/json"
	"fmt"
)

type V struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var b = []byte(`{"name": "foo", "age": 10}`)
var b2 = []byte(`{"age": 12}`)

func main() {
	var v V
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}
	fmt.Printf("before:%+v\n", v)

	// foo is still exists
	if err := json.Unmarshal(b2, &v); err != nil {
		panic(err)
	}
	fmt.Printf("after:%+v\n", v)
}

// unmarshal.
package main

import (
	"encoding/json"
	"fmt"
)

type JSON struct {
	String string   `json:"string"`
	Array  []string `json:"array"`
}

var str = []byte(`{
	"string": "hello world",
	"array": [
		"string 1",
		"string 2"
	]
}`)

var invalidStr = []byte(`{ string: "hello world" }`)

func main() {
	// valid
	var v JSON
	if err := json.Unmarshal(str, &v); err != nil {
		panic(err)
	}
	fmt.Printf("v:%+v\n", v)

	// invalid
	var v2 JSON
	if err := json.Unmarshal(invalidStr, &v2); err != nil {
		defer func() {
			fmt.Println(recover())
		}()
		panic(err)
	}
	fmt.Printf("v2:%+v\n", v)
}

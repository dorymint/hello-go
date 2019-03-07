// example for json
package main

import (
	"encoding/json"
	"fmt"
)

type JSON struct {
	String string   `json:"string"`
	Array  []string `json:"array"`
}

var contents = []byte(`{
	"string": "hello world",
	"array": [
		"string 1",
		"string 2"
	]
}`)

func main() {
	var v JSON
	if err := json.Unmarshal(contents, &v); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", v)
}

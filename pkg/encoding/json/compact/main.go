package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var j = []byte(`[
	{
		"name": "lily",
		"age": 11
	},
	{
		"name": "dory",
		"age": 12
	}
]`)

func main() {
	b := new(bytes.Buffer)
	err := json.Compact(b, j)
	if err != nil {
		panic(err)
	}

	fmt.Printf("json.Compact:\n")
	fmt.Printf("Before:%s\n", j)
	fmt.Printf("After:%s\n", b)
}

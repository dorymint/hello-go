package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "lily",
		Age:  11,
	}

	// Marshal
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(b))

	// with indent
	ib, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(ib))
}

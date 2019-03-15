// map.
package main

import (
	"encoding/json"
	"fmt"
)

var v struct {
	S  string             `json:"s"`
	M  map[string]*string `json:"m"`
	M2 map[string]*string `json:"m2"`
}

var b = []byte(`{
	"s": "string",
	"m2": {}
}`)

func main() {
	err := json.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON:%s\n", b)
	fmt.Printf("%+v\n\n", v)

	Print := func(m map[string]*string) {
		fmt.Printf("  is nil: %v\n", m == nil)
		fmt.Printf("  len   : %d\n\n", len(m))
	}

	fmt.Println("v.M:")
	Print(v.M)

	fmt.Println("v.M2:")
	Print(v.M2)
}

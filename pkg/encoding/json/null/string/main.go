// string pointer.
package main

import (
	"encoding/json"
	"fmt"
)

var M map[string]*string

const B = `{
	"foo": "bar",
	"fizz": null
}`

func main() {
	if err := json.Unmarshal([]byte(B), &M); err != nil {
		panic(err)
	}
	for key, val := range M {
		fmt.Printf("[%s]:%+v", key, val)
		if val != nil {
			fmt.Printf("  %s", *val)
		}
		fmt.Printf("\n\n")
	}
}

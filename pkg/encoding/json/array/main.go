// array.
package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Nodes []*Node

func (p *Nodes) MarshalIndent() {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func main() {
	var v Nodes

	fmt.Printf("before append:")
	v.MarshalIndent()

	v = append(v, &Node{
		Name: "foo",
		Age:  10,
	})

	fmt.Printf("after append:")
	v.MarshalIndent()
}

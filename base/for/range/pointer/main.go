// need care case of pointer used.
//
//	for _, v := range []*struct {}
//
package main

import "fmt"

type val struct {
	name string
	age  int
}

var input = []val{
	{name: "hello", age: 10},
	{name: "gopher", age: 11},
	{name: "world", age: 12},
}

const Expected = "gopher"

func miss(vs []val) (p *val) {
	for _, v := range vs {
		if v.name == Expected {
			p = &v
		}
	}
	return p
}

func correct(vs []val) (p *val) {
	for i, v := range vs {
		if v.name == Expected {
			p = &vs[i]
		}
	}
	return p
}

func main() {
	fmt.Printf("expected name: %q\n\n", Expected)

	fmt.Println("miss:")
	fmt.Printf("%+v\n\n", miss(input))

	fmt.Println("correct:")
	fmt.Printf("%+v\n\n", correct(input))
}

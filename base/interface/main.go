package main

import (
	"fmt"
)

type namePrinter interface {
	Print()
}

type base struct {
	name string
	id   int
}

func (b *base) Print() {
	fmt.Printf("%s %d\n", b.name, b.id)
}

type otherBase string

func (ob otherBase) Print() {
	fmt.Println(ob)
}

func main() {
	var np namePrinter

	np = &base{
		name: "base",
		id: 1,
	}
	np.Print()

	// not work, can accessed only declared functions
	//fmt.Println(np.name)

	// type cast
	switch t := np.(type) {
	case *base:
		fmt.Printf("base pointer: %s\n", t.name)
		t.name = "casting"
		t.Print()
	}

	np = otherBase("other base")
	np.Print()
}

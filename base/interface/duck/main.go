package main

import (
	"fmt"
)

type duck interface {
	SayHello() string
}

type cat struct {
	name string
	age  int
}

func (c *cat) SayHello() string {
	return "hello from cat"
}

type dog int

func (d dog) SayHello() string {
	return "hello from dog"
}

func main() {
	var d duck
	d = &cat{}
	fmt.Println(d.SayHello())
	d = dog(0)
	fmt.Println(d.SayHello())
}

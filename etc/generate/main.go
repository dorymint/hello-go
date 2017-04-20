package main

import (
	"fmt"
)


// syntax={//go:generate [command] [args]}
// not allow //<space>go:generate

//go:generate echo "hello generate"
const msg = `go:generate test

syntax = {//go:generate [command] [args]}
not allow = {// go:generate [command] [args]}
example = {//go:generate echo "hello generate"}

	try it
RUN:
	go generate ./
OUT:
	hello generate
`
func main() {
	fmt.Println(msg)
}

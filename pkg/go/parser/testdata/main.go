// (,,ﾟДﾟ)
package main

import (
	"flag"
	"fmt"
)

/* (,,ﾟДﾟ)

// (,,ﾟДﾟ)

(,,ﾟДﾟ) */

var nyan = flag.String("n", "(,,ﾟДﾟ)", "giko nyan")

type Struct struct {
	Name string
}

func hi() { print("hi") }

func main() {
	var at = string("hello")
	flag.Parse()
	fmt.Println(*nyan)
	fmt.Println(at)
}

// (,,ﾟДﾟ)

// IndexRune.
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "192.168.1.1:8080"
	fmt.Println("str:")
	fmt.Printf("\t%s\n", str)
	fmt.Println("trim ports by index: str[:strings.IndexRune(str, ':')]")
	fmt.Printf("\t%+v\n", str[:strings.IndexRune(str, ':')])
}

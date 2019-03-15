// equal.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte(`hello world`)
	fmt.Println(bytes.Equal(b, b)) // true

	bb := []byte(`not`)
	fmt.Println(bytes.Equal(b, bb)) //false

	bbb := []byte(`hello world`)
	fmt.Println(bytes.Equal(b, bbb)) // true
}

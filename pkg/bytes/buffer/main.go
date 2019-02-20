package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Println("string:", b.String(), "bytes:", b.Bytes())

	echo := func() func() []byte {
		b := []byte("world\n")
		return func() []byte {
			b = append(b, b...)
			return b
		}
	}()

	fmt.Println("push \"world\\n\"")
	for i := 0; i < 10; i++ {
		b.Write(echo())
	}
	fmt.Println("count of [world]", bytes.Count(b.Bytes(), []byte("world")))
	b.Reset()
	fmt.Println("after reset", b.String(), b.Bytes())
}

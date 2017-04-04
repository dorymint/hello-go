package main

import (
	"bytes"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))

	println(b.String(), b.Bytes())
	echo := func() func() []byte {
		b := []byte("lily\n")
		return func() []byte {
			b = append(b, b...)
			return b
		}
	}()
	for i := 0; i < 10; i++ {
		b.Write(echo())
	}
	println("count of [lily]", bytes.Count(b.Bytes(), []byte("lily")), b.Bytes())
	b.Reset()
	println(b.String(), b.Bytes())
}

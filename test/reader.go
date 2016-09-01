package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)


// infinityReader Aを出し続けるReaderをmethodに持つ
type infinityReader struct {}

func (p *infinityReader) Error() string {
	return fmt.Sprintf("eol")
}

func (p infinityReader) Read(b []byte) (int, error) {
	n := 0
	for i := range(b) {
		b[i] = byte('A')
		n++
	// TODO:break
	}
	return n, nil
}

func main() {
	fmt.Printf("%v\n", byte('A'))

	r := infinityReader{}

	// gift io.Reader
	reader.Validate(r)
}

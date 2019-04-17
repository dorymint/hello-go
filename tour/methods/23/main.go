// Exercise: rot13Reader.
package main

import (
	"io"
	"os"
	"strings"
)

var r13m = func() map[rune]rune {
	m := make(map[rune]rune)
	a, n := 'a', 'n'
	A, N := 'A', 'N'
	for a != 'n' {
		m[a], m[n] = n, a
		a++
		n++
		m[A], m[N] = N, A
		A++
		N++
	}
	return m
}()

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := range b {
		r, ok := r13m[rune(b[i])]
		if ok {
			b[i] = byte(r)
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

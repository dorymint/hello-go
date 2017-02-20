// check Read() method

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	b := make([]byte, 1<<1)
	fmt.Println("b:len,cap", len(b), cap(b))

	// バッファの容量不足時
	// reader 内部の index は途中まで動く
	src := []byte("hello world")
	r := bytes.NewReader(src)
	fmt.Println("r.len:", r.Len(), "src.len:", len(src))

	// 記録されている index から先を読み取る
	n, err := r.Read(b)
	fmt.Println("n,err", n, err)
	fmt.Println("r.len", r.Len())
	fmt.Println(string(b))

	for {
		n, err := r.Read(b)
		fmt.Println(string(b), err)
		// 終端はEOFが返る
		if err == io.EOF { fmt.Println("EOF!!") }
		if n == 0 {
			break
		}
	}
	// もう一度読むとbyteは入らずEOFが返る
	onemore := make([]byte, 1<<2)
	_, err2 := r.Read(onemore)
	fmt.Println("one more", onemore, err2) // err2=EOF


	// Len はまだ読まれてない byte の数
	// Size は reader が読んでるオリジナルの len
	fmt.Println(r.Len(), r.Size())
	fmt.Println(r)
	// r = {[]byte(...), 11(current reading index), -1(index of previous rune:or <0)
	seeknum, err := r.Seek(0, 0)
	fmt.Println(r, seeknum, err)
	b2 := make([]byte, 1<<4)
	r.Read(b2)
	for i, j := range b2 {
		if j == byte(0x0) {
			b2[i] = ' '
		}
	}
	fmt.Println(r)
	fmt.Println(string(b2), r.Len())
}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// DONE:make rot13Reader.Read()
func (rot *rot13Reader) Read(p []byte) (int, error) {
	// 注意:前に r 忘れて無限再帰してスタックオーバーフローした
	slice := make([]byte, 128, 256)
	n, err := rot.r.Read(slice)
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "不明なエラーが発生しました")
		os.Exit(1)
	}

	// 復号化してpに詰める
	rotmap := newRotmap()
	m := int(0)
	for i := range(p) {
		if i < n {
			p[i] = rotmap[slice[i]]
			m++
		} else {
			return m, fmt.Errorf("slice[%v] length caution", i)
		}
	}
	return m, nil
}

// せっかくなのでmap使ってみる
func newRotmap() map[byte]byte {
	newMap := make(map[byte]byte)
	// ASC2
	// 単純なメモ化だけど書き方が...もう少し何とかしたい
	for a,A, n,N := byte('a'),byte('A'), byte('n'),byte('N'); a <= byte('m'); a,A, n,N = a+1,A+1, n+1,N+1 {
		// 少文字
		newMap[a] = n
		newMap[n] = a

		// 大文字
		newMap[A] = N
		newMap[N] = A
	}
	newMap[byte(' ')] = byte(' ')
	newMap[byte('\n')] = byte('\n')

	return newMap
}

func main() {
	fmt.Println("test")
	// io.Readerをsで受ける
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// io.Reader s を rot13Reader.r に含んだ rot13Reader r を作る
	r := rot13Reader{s}
	// io.Copy で標準出力へ r の内容を吐き出す
	io.Copy(os.Stdout, &r)
}


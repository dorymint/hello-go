package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

// ファイルストリームの取り扱い
// go的に言えばデータストリーム?

/* ioパッケージ */
/* io.Readerインターフェースを既定 */
// ctrl-kすれば見れるけど下のメソッドが実装されてる
// func (r *Reader) Read(b []byte) (n int, err error)

// 'A'を無限に出力するReaderの実装
type InfinityAReader struct {}
// error型はnil返せばいいのでこのError()はなくても一応大丈夫
func (p *InfinityAReader) Error() string {
	return fmt.Sprintf("eol")
}
// 'A'を無限に出力するりーだー,byteスライスを渡すと全て'A'で埋まる
func (infA InfinityAReader) Read(p []byte) (int, error) {
	n := int(0)
	for i := range(p) {
		p[i] = byte('A')
		n++
	}
	// 正常系
	return n, nil
}

func main() {

	// std::sstream 的な
	r := strings.NewReader("Hello, Reader!!")
	fmt.Printf("%T, %p, %v\n", r, r, r)

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		// NOTE:
		// Read()は自身の内容を引き出し、引数に渡したbyte配列に入れようとする
		// 戻り値は詰めたlength(int)とerr
		// byte配列のlengthが足りなければ終了する
		// データを詰めてeofに達すると正常終了する
		// streamがeofに達していた時にerrが返る

		fmt.Printf("n = %v err = %v b = %v\n", n,  err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { break }
	}


	// 自作いんふぃにてぃAりーだーのテスト
	reader.Validate(InfinityAReader{})
}


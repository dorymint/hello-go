// scriptencoding utf-8

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func main() {

	// hello world
	fmt.Printf("hello go\n")
	str := "hello"
	fmt.Printf("%s\n", str)

	// パッケージの関数とか分からなければgodocする

	fmt.Println("welcome to the go tuto!")

	// packageの追加と外部関数の使用
	fmt.Println("The time is", time.Now())
	fmt.Println("go run test")

	// rand乱数生成, seadが固定なので実行結果も固定
	fmt.Println("My favarid number is", rand.Intn(10))

	// べき乗計算の計算結果を表示,精度はgodoc参照
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))

	// packagename.<大文字で始まる名前> は外部から参照できるようにexportされている
	fmt.Println(math.Pi)

	// func
	fmt.Printf("call add(1, 10) = %d\n" , add(1, 10))
	fmt.Println(add(1, 10) , "Println delimiter is ','" )

}

// :nnoremap <c-@>gorun :!go run % > gorunlog.txt<nl>

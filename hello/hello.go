// set fileencoding=utf-8

// NOTE:パッケージの関数とか分からなければgodocする

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 変数の型は前方ではなく後ろで明示できる
func add(x int, y int) int {
	return x + y
}

// NOTE:golintはキャメルケースを推奨してるっぽい?スネークケースはgolintで注意される
// 引数の型は最後の引数と同じなら省略できる
func add_another(x, y int) int {
	return x + y
}

// hello world
func hello_go() {

	// hello world
	fmt.Printf("hello go\n")
	str := "hello"
	fmt.Printf("%s\n", str)
	fmt.Println()

	fmt.Println("welcome to the go tuto!")
	fmt.Println()

	// packageの追加と外部関数の使用
	fmt.Println("The time is", time.Now())
	fmt.Println("go run test")
	fmt.Println("Println delimiter is ','", time.Now())
	fmt.Println()

	// rand乱数生成, seadが固定なので実行結果も固定
	fmt.Println("My favarid number is", rand.Intn(10))
	fmt.Println()

	// べき乗計算の計算結果を表示,精度はgodoc参照
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println()

	// packagename.<大文字で始まる名前> は外部から参照できるようにexportされている
	fmt.Println("math.Pi is ", math.Pi)
	fmt.Println()

	// func
	fmt.Printf("call add(1, 10) = %d\n", add(1, 10))
	fmt.Println("add_another(1, 20) is ", add_another(1, 20))
	fmt.Println()

	// 関数定義の上下は問題無いっぽい
	fmt.Println("string swap")
	x, y := "hello", "world"
	fmt.Printf("%s\n", x+y)
	x, y = swap(x, y)
	fmt.Printf("%s\n", x+y)
	fmt.Println()

	// naked return
	fmt.Println(split(17))
	i, z := split(27)
	fmt.Println(i, z)
	fmt.Println()

	return
}

// 引数を入れ替えて返す,goは複数の戻り値を持つことができる
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値の変数名
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// naked return
	return
}

// variables, 全てbool型になる
var c, python, java bool

func variable() {
	var i int
	fmt.Println(i, c, python, java)
	return
}

// var initalizer
var i, j int = 1, 2

func var_ini() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	return
}

// 暗黙の型宣言
var test = "test"

func short_var() {
	fmt.Println(test)

	// := は関数内で使える暗黙の代入文
	test := "short variable declaration"
	fmt.Println(test)

	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)

	return
}

func main() {
	hello_go()
	fmt.Println()

	variable()
	fmt.Println()

	var_ini()
	fmt.Println()

	short_var()
	fmt.Println()

	return
}

// vim
// NOTE:レジスタを使ってコマンドを貼り付ける時は<C-r>"で実現できる
// :nnoremap <c-@>gorun :!go run % > gorunlog.txt 2>&1<nl>

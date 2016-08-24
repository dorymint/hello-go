package main

import (
	"container/list"
	"fmt"
	"math/cmplx"
)

/// {{{ variable types
// from gotuto
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var (
	boolean bool   = true
	str     string = "this string"

	// 型はコンストラクタを持ってるっぽい
	integer   int   = int(10)
	integer8  int8  = int8(8)
	integer32 int32 = int32(32)
	integer64 int64 = int64(64)

	uinteger   uint   = 10
	uinteger8  uint8  = 8
	uinteger32 uint32 = 32
	uinteger64 uint64 = 64

	ptr     uintptr
	bit8    byte = 100
	unicode rune = 'a'

	floating32 float32 = 10.123456789
	floating64 float64 = 100.123456789

	com64  complex64  = complex64(1)
	com128 complex128 = cmplx.Sqrt(-5 + 12i)
)

// -----| goの組み込み型 |-----
// bool
// string
// int int8 int32 int64
// uint uint8 uint32 uint64 uintptr
// byte == uint8
// rune == int32 // NOTE:ユニコードのコードポイント
// float32 float64
// complex64 complex128

// NOTE:int uint uintptrはシステムに合わせて 32bit or 64bit が使われる

// 変数の初期値は決まっていて代入がなければ初期化されるらしい
func zeroValus() {
	var (
		i int
		f float64
		b bool
		s string
	)
	fmt.Printf("%v %v %v %q", i, f, b, s)
	// %qはダブルクォート方式のエスケープ
	fmt.Println()
}

// goの変数の型を確認
func basicTypes() {
	// %v --デフォルトフォーマット?
	// autoで型を出すことを言っているのか?
	// %T --goスタイルのフォーマット?
	// TODO:ちょっと何を言ってるか理解してないので後で調べる
	// 調べてないけど多分%Tがタイプを表示して%vが元の型で値を表示するだけ

	// format
	const f = "%T(%v)\n"

	// from gotuto
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// variable list view
	varlist := list.New()

	varlist.PushBack(boolean)
	varlist.PushBack(str)

	varlist.PushBack(integer)
	varlist.PushBack(integer8)
	varlist.PushBack(integer32)
	varlist.PushBack(integer64)

	varlist.PushBack(uinteger)
	varlist.PushBack(uinteger8)
	varlist.PushBack(uinteger32)
	varlist.PushBack(uinteger64)

	varlist.PushBack(ptr)
	varlist.PushBack(bit8)
	varlist.PushBack(unicode)

	varlist.PushBack(floating32)
	varlist.PushBack(floating64)

	varlist.PushBack(com64)
	varlist.PushBack(com128)

	// display to variable
	for e := varlist.Front(); e != nil; e = e.Next() {
		tmp := e.Value
		fmt.Printf(f, tmp, tmp)
	}
	fmt.Println()

	// another format
	const f2 = "%T(%v) %#v\n"
	for e := varlist.Front(); e != nil; e = e.Next() {
		tmp := e.Value
		fmt.Printf(f2, tmp, tmp, tmp)
		// #付きと付いてないフォーマットの違いがいまいちつかめてない
	}
	fmt.Println()

	// 変数の初期値
	zeroValus()
	fmt.Println()

	return
}
/// variable types }}}

/// array {{{
// display slice
func dispSlice(slice []int) {
	fmt.Println(slice)
	for _, x := range slice{
		fmt.Printf("%c\n", x)
	}
	fmt.Println()
	return
}

func array() {
	// array
	fmt.Println("配列")
	var a [4]int
	fmt.Println(a)
	a[0] = 1
	fmt.Println(a)
	i := a[0]
	fmt.Println(i)
	fmt.Printf("%v (%T)\n", a, a)
	fmt.Printf("%v %v\n", a[0], a[1])
	// 範囲外アクセスはコンパイルエラー
	// fmt.Printf("%v %v", a[4], a[-1])

	// 整数と文字
	fmt.Println("整数と文字")
	a[0] = 97
	a[1] = 98
	a[2] = 99
	a[3] = 100
	fmt.Printf("%s\n", a) // aは配列全体を指していて、先頭要素へのポインタではない
	fmt.Printf("%q\n", a)
	fmt.Printf("%x\n", a) // 2文字の16進数表記
	fmt.Printf("%p\n", a)
	for _, x := range a {
		fmt.Printf("%c\n", x)
	}
	fmt.Println()


	// スライス
	fmt.Println("スライス")
	letter := []int{97, 98, 99, 100}
	dispSlice(letter)
	// 範囲を指定して削る 1<= ... >3  // [3]の要素から含まれない
	fmt.Println("スライスの範囲指定")
	slice := letter[1:3]
	dispSlice(slice)
	// sliceの省略
	fmt.Println("スライスの省略")
	slice_1 := letter[:] // [0] <= [3]
	dispSlice(slice_1)
	slice_2 := letter[1:] // [1]<
	dispSlice(slice_2)
	slice_3 := letter[:3] // >[3]
	dispSlice(slice_3j
}
/// array }}}

func main() {

	basicTypes()
	array()

	return
}

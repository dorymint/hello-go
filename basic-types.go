
package main

import (
	"container/list"
	"fmt"
	"math/cmplx"
)

// from gotuto
var (
	ToBe	bool		= false
	MaxInt	uint64		= 1<<64 -1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
)


var (
	boolean		bool	=	true
	str			string	=	"this string"

	// 型はコンストラクタを持ってるっぽい
	integer		int		=	int(10)
	integer8	int8	=	int8(8)
	integer32	int32	=	int32(32)
	integer64	int64	=	int64(64)

	uinteger	uint	=	10
	uinteger8	uint8	=	8
	uinteger32	uint32	=	32
	uinteger64	uint64	=	64

	ptr			uintptr
	bit8		byte	=	100
	unicode		rune	=	'a'

	floating32	float32	=	10.123456789
	floating64	float64	=	100.123456789

	com64		complex64	=	complex64(1)
	com128		complex128	=	cmplx.Sqrt(-5 + 12i)
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

func main() {

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


	return
}



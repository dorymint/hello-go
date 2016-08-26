package main

import (
	"fmt"
)

func array_1() {
	fmt.Println("array_1")
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
	fmt.Printf("%s\n", a)
	fmt.Printf("%q\n", a)

	// 配列の長さは変えられないが、スライスを使って切り取りコピーできる
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	return
}

func slice_1() {
	fmt.Println("slice_1")

	// スライスは配列へのポインタを内部に持つクラスのようなものと理解している(若干理解が怪しい)
	// 配列への capacity, lenge, pointer, を内部に持つ

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)


	// スライスリテラル、配列と似た形で宣言、要素数を明示していない
	sliceLiteral := []string{ "a", "b", "c", "d" }
	fmt.Println(sliceLiteral, "\n")

	// スライスはデータの実体ではなくポインタを持つため、同じ実体を共有いしている変数に影響する
	fmt.Println("共有しないスライス")
	a := []string{"hello", "world", "startup"}
	fmt.Println("a = ", a)
	b := []string{"cpp", "python", "golang"}
	fmt.Println("b = ", b)

	fmt.Println("a = appjnd(a, b...)\n")
	a = append(a, b...) // スライスの追加関数,新しく実体を作ってる？
	fmt.Println("a = ", a, "\n")

	fmt.Println("a[4] = string(\"ruby\")")
	a[4] = string("ruby")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b, "\n") // aと共有していない

	// 共有するスライス
	fmt.Println("共有するスライス")
	p := a[:]
	fmt.Println("p = ", p)
	p[5] = string("lua")
	fmt.Println("a = ", a)

	// append()すると実体が別れる？
	// このあたり若干理解が怪しい

	return
}

func main() {

	array_1()
	fmt.Println()
	slice_1()

}

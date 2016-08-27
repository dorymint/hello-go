package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"strings"
)

// 配列
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


// slice startup
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

// array to slice
func slice_2() {
	fmt.Println("slice_2")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"	// pointer like
	fmt.Println(a, b)
	fmt.Println(names)
	return
}

// slice
func slice_3() {
	fmt.Println("slice_3")
	// 配列と似てるけど作ってるのはsliceで参照を持つオブジェクトっぽい

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	return
}

// slice literal
func slice_4() {
	// sliceをslice
	fmt.Println("slice_4")
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	return
}

// slice pointer, lenge, capacity
func slice_5() {
	fmt.Println("slice_5")
	// sliceはpointer lenge capacity を持つ
	printSlice := func (s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
		return
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:] // sliceの頭を切ると、切ったレンジのキャパが落ちてる
	printSlice(s)


	// nil slice
	fmt.Println("\nnil slice")
	var nilSlice []int // 宣言だけ,初期値はnil
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	return
}

// make
func slice_6() {
	fmt.Println("slice_6")
	// sliceを作るmake関数

	printSlice := func (s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n", s,  len(x), cap(x), x)
		return
	}

	// make() parameter is type, lenge(optional), cap(optional)
	a := make([]int, 5) // lengeを指定しないとcap分確保して初期化される
	printSlice("a", a)

	b := make([]int, 0,  5) // capのみ確保してる
	printSlice("b", b)

	c := b[:2] // 頭は切ってもcapは削られない
	printSlice("c", c)

	d := c[2:5] // 尻尾を着るとcapも削られるっぽい
	printSlice("d", d)


	return
}

// slice of slice, join
func slice_7() {
	fmt.Println("slice_7")

	// 二次元配列
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
		// joinは文字列配列を受けて要素を指定した文字列で区切ってくっつけた文字列を返す
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
		// 文字列配列が返る
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
		// separatorは省略できない
	}
	// 二次元配列の表示の時にjoinを忘れると配列のまま返るので注意する

	return
}

// append
func slice_8() {
	fmt.Println("slice_8")

	printSlice := func ( x []int) {
		fmt.Printf("address=%p len=%d cap=%d %v\n", &x, len(x), cap(x), x)
		return
	}

	// append() parameter, slice, type-valus...
	var s []int
	printSlice(s)

	// capacityが足りなければ新しく確保した領域で拡張されたsliceを作る
	// 元のスライスとは異なるsliceの実体を返すっぽい?
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)


	// capacityを多めにとったslice
	fmt.Println("\ncapacityを多めにとったスライス")
	s_2 := make([]int, 0, 10)
	printSlice(s_2)

	s_2 = append(s_2, 0)
	printSlice(s_2)

	s_2 = append(s_2, 1)
	printSlice(s_2)

	s_2 = append(s_2, 2, 3, 4)
	printSlice(s_2)


	// どちらもアドレスが変わってる、capacity多めにとっても
	// appendは新しいsliceを作って返してる?

	return
}

// range
func sliceRange() {
	// sliceに対するrangeは要素数のカウンタと要素のコピーを返す

	fmt.Println("sliceRange")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256}
	for i, v := range(pow) {
		fmt.Printf("2**%d = %d\n", i, v)
	}


	// _(アンダースコア)を変数名にするとデータを使わないことを明示できる
	// 返ってきた変数を使わなくてもエラーにならない
	pow_2 := make([]int, 10)
	for i := range pow_2 {
		pow_2[i] = 1 << uint(i) // == 2**i
	}
	for _, pow_2 := range(pow_2) {
		fmt.Printf("%d\n", pow_2)
	}
	return
}


func slicePic() {
	fmt.Println("slicePic")

	Pic := func (dx, dy int) [][]uint8 {
		slice := make([][]uint8, dy)
		elm := make([]uint8, dx)

		for i, _ := range(slice) {
			for _, e := range(elm) {
				slice[i] = append(slice[i], e)
			}
		}
		// 取り敢えず初期化だけ
		return slice
	}

	// 表示はブルースケール全部0の初期値で青一色
	pic.Show(Pic)

	return
}

func main() {

	array_1()
	fmt.Println()
	slice_1()
	fmt.Println()
	slice_2()
	fmt.Println()
	slice_3()
	fmt.Println()
	slice_4()
	fmt.Println()
	slice_5()
	fmt.Println()
	slice_6()
	fmt.Println()
	slice_7()
	fmt.Println()
	slice_8()
	fmt.Println()
	sliceRange()
	fmt.Println()
	slicePic()
	fmt.Println()




}

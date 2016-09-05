package main

import (
	"fmt"
)

type I interface {
	M()
}
type emptyI interface{}
// 空のインターフェース型はメソッドの定義が無い型も受け入れる

type T struct {
	S string
}
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>\n")
		return
	}
	fmt.Println(t.S, "\n")
	// tがnilだとここでエラーになるのでifで処理してる
}

type myInt int
func (i myInt) M() {
	fmt.Println(i, "\n")
}


// NOTE:interfaceは値への参照と型情報のタプル

func main() {

	/* interface */
	var i I
	var general emptyI
	//i.M() // panic nil.M() インターフェース型の初期値はnil
	fmt.Printf("T=%T v=%v i\n", i, i)

	// 汎用インターフェース型
	fmt.Printf("\ngeneral\nT=%T v=%v\n", general, general)
	general = 42
	fmt.Printf("T=%T v=%v\n", general, general)
	general = "hello"
	fmt.Printf("T=%T v=%v\n\n", general, general)

	fmt.Println("*T")
	var t *T
	i = t
	describe(i)
	i.M()
	// iに値の実体がなくても型情報だけで実装されている関数を呼べるっぽい
	// M()でnilの処理を入れているためエラーにならない

	i = &T{"hello"}
	describe(i)
	i.M()

	fmt.Println("myInt")
	i = myInt(10)
	describe(i)
	i.M()


	/* インターフェースに格納した値を取り出す */
	//var intValu int = i // error

	// i.(T) 型アサーションと呼ばれる記法で取り出せる
	fmt.Println("i.(myInt) = ", i.(myInt)) // myInt(10)が取り出せる

	// intへの代入
	var toInt int = int(i.(myInt))
	fmt.Println(toInt, "\n")

	// 上述した記法は完全では無い、アサーションが失敗した場合は戻り値が2つある
	toT, state := i.(*T)
	fmt.Println(toT, state, "\n")
	// 失敗した時は型の初期値とfalseが返る

	// 戻り値を1つの変数で受けてアサーションが失敗した場合
	//err2 := i.(*T) // panic 戻り値の1つを_で捨てればpanicは回避できる
	t2, _ := i.(*T)
	fmt.Println("t2 = ", t2, "\n")



	/* switchとinterface */
	fmt.Println("switch and interface")
	// x := i.(type) この記法は代入では使えない
	switch x := i.(type) {
	case myInt:
		fmt.Printf("%T %v %p\n\n", x, x, x)
	default:
		fmt.Println("default")
	}
	// xがswitch用のタプルの様に見える
	// caseで比較されるのは型
	// ブロックで使われるxは値への参照
	do := func(I interface{}) {
		switch v := I.(type) {
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	do("hello")
	do(21)
	do(true)



	/* panic test */
	fmt.Println("\npanic test")
	var i2 I
	describe(i2)
	//i2.M() // nil.M() panic

	var mi *myInt
	i2 = mi
	describe(i2)
	//i2.M() // panic.
	// iはi=miの時点でnilと*myIntのタプルを持つ
	// i.M()でnilをmyIntへ渡そうとしてしまうのでパニックになる

}

func describe(i I) {
	f := "i = %p valueP = %p\n(%v, %T)\n"
	fmt.Printf(f, &i, i, i, i)
	// %p = &i インターフェース型タプルへのアドレス
	// %p = i  値へのアドレス
}



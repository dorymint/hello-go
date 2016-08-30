package main

import (
	"fmt"
	"math"
)


type I interface {
	M()
}
type T struct {
	S string
}
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
	
	// interface型が持ってるのは値と型のtuple
	// 実体への参照を直に持たず、tupleを持ってる
	// I型の変数の宣言だけの初期値は<nil>
	var declOnly I
	fmt.Printf("%T\n", declOnly)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
	// %T i は main.I 型が返りそうだけど参照先の型情報が返る
}

package main

import (
	"fmt"
	"math"
)

/*
	interfaceはメソッドのシグネチャの集まり
	クラスメソッドの宣言だけ集めてる
*/


// type "name" interface {} でinterface型の変数を宣言している
type Abser interface {
	Abs() float64
}

// 変数レシーバ
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ポインタレシーバ
type Vertex struct { X, Y float64 }
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {

	format := "%T %v\n"

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	fmt.Printf(format, a, a)
	fmt.Printf(format, f, f)
	fmt.Printf(format, v, v)

	a =  f; fmt.Printf(format, a, a)
	a = &v; fmt.Printf(format, a, a)


	// vは Abser で定義したメソッドの実装を持っていないためerror
	// a =  v; fmt.Printf(format, a, a)

	// v. は (&v). と解釈される
	// 変数レシーバが実装されているように動くのでinterfaceの代入で混乱しそう
	fmt.Println(v.Abs())

	// 代入された内容はそのまま使える
	fmt.Println(a.Abs())


	// 汎用interface使ってジェネリクスっぽいの
	type inter interface{}
	puts := func (T inter) {
		format := "%T %v\n"
		fmt.Printf(format, T, T)
	}
	puts(v)

	main2()

}


type I interface { M() }
type T struct { S string }

// interfaceを実装することを明示する必要は無い
func (t T) M() { fmt.Println(t.S) }

func main2() {
	var i I = T{"hello"}
	i.M()
}


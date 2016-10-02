package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/// メソッドをポインタレシーバとして省略する

/// ポインタレシーバ {{{
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// (v *Vertex) を上で定義しているため、メソッドの再定義でerror吐く
// func (v Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }
func method_1() {
	fmt.Println("method_1")

	v := Vertex{3, 4}
	v.Scale(2) // (&v). を省略してv. で行ける、goが(&v)の省略と解釈するらしい
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := &v
	p.Scale(3) // こちらはsrcの見た目通りアドレスを渡している
	ScaleFunc(p, 3)
	fmt.Println(p)

	fmt.Println(v, p)
}
/// }}}


/// 変数レシーバ {{{
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func method_2() {
	fmt.Println("method_2")

	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 値に実装されているメソッドなので見た目通り
	fmt.Println(AbsFunc(v))

	p := &v
	fmt.Println(p.Abs()) // p == (*p) と解釈される、見た目とは異なるが便利
	fmt.Println(AbsFunc(*p))
}
/// }}}

/// ポインタレシーバのメモリ効率
func (v * Vertex) Abs_p() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func memoryView() {

	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs_p())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs_p())

}

/*
	型に実装するレシーバはポインタか変数どちらか一方に統一するべき
*/

func main(){

	method_1()
	fmt.Println()
	method_2()
	fmt.Println()
	memoryView()
	fmt.Println()

}

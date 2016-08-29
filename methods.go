package main


import (
	"fmt"
	"math"
)

// 型に定義する関数のtutorial


/// structに対するメソッドの定義 {{{
type Vertex struct {
	X, Y float64
}
// (v Vertex) がレシーバ
func(v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// レシーバを使わずに同じ機能
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
/// }}}

// structでなくてもメソッド宣言できる
type myFloat float64
func(f myFloat) Abs() float64 {
	// goに3項演算子は無い?
	if f < 0 {
		return float64(-f) // -fで符号逆転できる,直感的わかりやすい
	} else {
		return float64(f)
	}
}


func main() {

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))


	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Printf("%T %v %v\n", f, f, f.Abs())
		// %T = main.myFloat 型
		// "%v", f で値が返るのはなぜか? fはmyFloat型の値のはずだけど%vで普通に表示される
		// 既定の型を見て%vによる値の表示方法を決めてるっぽい?

}

package main


import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}


func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// ポインタレシーバ
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// Scaleメソッドを関数で書きなおす
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) ScaleX(f float64) { v.X = v.X * f }
func (v *Vertex) ScaleY(f float64) { v.Y = v.Y * f }

func main() {

	// ポインタレシーバ
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	// ポインタレシーバ
	fmt.Println(v)
	v.ScaleX(10)
	fmt.Println(v)
	v.ScaleY(10)
	fmt.Println(v)

	// 関数 &を付けないとアドレスが渡らない、値を渡すことになる
	Scale(&v, 0.1)
	fmt.Println(v)
}

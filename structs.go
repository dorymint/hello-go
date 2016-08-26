package main

import (
	"fmt"
)

// goの構造体
type Vertex struct {
	X int
	Y int
}

func struct_1() {
	fmt.Println("struct_1")

	fmt.Println(Vertex{1, 2})

	// 構造体の初期化は中括弧で括るっぽい,listで渡す
	v:= Vertex{5, 2}
	fmt.Println(v)

	// use pointer
	p := &v
	p.X = 1e9 // ->とか(*p).とか使わなくてもいいらしい
	fmt.Println(v)

	// direct access
	v.Y = 1e5
	fmt.Println(v)
	return
}

func struct_2() {
	fmt.Println("struct_2")

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}	// Vertex.Xだけを初期化している
		v3 = Vertex{}		// 初期値
		p  = &Vertex{1, 2}	// structへのポインタ、実体はあるが変数には代入していない状態?
	)

	fmt.Println(v1, v2, v3, p)
	return
}

func main() {

	struct_1()
	fmt.Println()
	struct_2()

}

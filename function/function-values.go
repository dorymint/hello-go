package main

import (
	"fmt"
	"math"
)

func funcValue_1() {
	fmt.Println("funcValue_1")

	compute := func(fn func(float64, float64) float64) float64 {
		return fn(3, 4)
	}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// function closure
func funcValue_2() {
	fmt.Println("funcValue_2")

	// func (引数なし) 戻り値 func(int) int{} のオブジェクト
	addre := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	// クロージャをを返すaddreを呼んで返ったクロージャを変数に入れる
	// sumはクロージャの外で宣言されているがクロージャ内から参照できる
	// 変数は使われる限り残るのでsumの実体も残り、使われなくなるまで値を保持する
	// sumは各実体ごとに異なるため値は干渉しない
	// sumが横に干渉しないstaticっぽい動きに見える
	pos, neg := addre(), addre()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	return
}

// fibonacci
func fibo() {
	fibonacci := func() func() int {
		// memo まずは素直に
		result := int(0)
		next := int(1)
		var tmp int

		walker := func() {
			tmp = result + next
			result = next
			next = tmp
		}
		// 他にも幾つかやり方があると思うけど一応これで
		return func() int {
			defer walker()
			return result
		}
	}
	// defer walker() は下でもいける
	// defer func() { tmp = result + next; result = next; next = tmp; }()


	// view loop
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	return
}

func main() {

	funcValue_1()
	fmt.Println()

	funcValue_2()
	fmt.Println()

	fibo()
	fmt.Println()

}

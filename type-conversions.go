package main

import (
	"fmt"
	"math"
)

func typeInference() {

	// int
	v := 42
	fmt.Printf("v is of type %T\n", v)

	// float
	f := 42.195
	fmt.Printf("v is of type %T\n", f)

	// complex128
	g := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", g)

	return
}


/// constnts {{{
func constants() {

	// 定数は型推論できない
	const Pi = 3.14
	const Str = "world"
	fmt.Println("hello ", Str)
	fmt.Println("happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	return
}

// 数値の定数は精度が高い
const (
	Big = 1 << 100
	Small = Big >> 99
)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func numericConstants(){

	fmt.Println("const values")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

//	fmt.Println("is valus overflows")
//	var (
//		big = 1 << 100
//		small = big >> 99
//	)
//	fmt.Println(needInt(small))
//	fmt.Println(needFloat(small))
//	fmt.Println(needFloat(big))

	return
}
/// constants }}}

func main() {

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	func() {
		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)
		// 型の変換は明示しなければコンパイルエラー

		fmt.Println(f, i, u)
	}()

	typeInference()
	constants()
	numericConstants()
}

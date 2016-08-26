package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/// loop tutorial {{{
func loop() {
	f := "result %v\n\n"

	// for loop
	fmt.Println("for loop")
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += 1
	}
	fmt.Printf(f, sum)


	// for 初期化と後処理は省略可
	fmt.Println("省略形")
	sum = 1
	for ; sum < 100; {
		sum += sum
		if sum % 2 == 0 {
			fmt.Println(sum)
		}
	}
	fmt.Printf(f, sum)


	// while ...forのセミコロンも省略できる
	fmt.Println("while loop")
	sum = 1
	for sum < 1000 {
		sum += sum
		if sum % 8  == 0 {
			fmt.Println(sum)
		}
	}
	fmt.Printf(f, sum)

	// forever
	fmt.Println("infinite loop")
	for {
		sum -= 10
		fmt.Println("hello!! " , sum)
		if sum < 0 { break }
	}

	fmt.Println()
	return
}
/// loop tutorial end }}}


/// if tutorial {{{

func sqrt(x float64) string {
	if x < 0 {
			return sqrt(-x) + "i"
		}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// forのように条件の前にifスコープ内で有効な文を書ける
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // disp 1
	}
	// can't use v here, though
	return lim // disp 3
}

// ニュートン法を使った収束
func newton(x float64) float64 {
	// TODO
	x2 := float64(x)
	z := float64(5)
	i := 0
	for ; i < 10000; i++ {
		x2 = z - (z * z - x) / (2*z * x)

		if math.Abs(x2 - z) < 0.0000001 { break }

		z = x2

	}
	fmt.Println("loop cout = ", i)
	return x2
}

func iftuto() {
	// if statement
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println()

	fmt.Println(
		pow(3, 2, 10), // disp 2
		pow(3, 3, 10), // disp 3
	)
	fmt.Println()

	// newton
	fmt.Println("newton")
	for i := 1; i < 10; i++ {
		fmt.Println("i = ", i)
		fmt.Println(sqrt(float64(i)))
		fmt.Println(newton(float64(i)))
		fmt.Println()
	}
	fmt.Println()

	return
}
/// if tutorial end }}}

/// switch {{{
func switch_1() {
	// goのswitchはデフォルトでbreakする
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	return
}

func switch_2() {
	// switchのケースに式を使う
	fmt.Println("When's Saturday?")
	tody := time.Now().Weekday()
	switch time.Saturday {
	case tody + 0:
		fmt.Println("Tody.")
	case tody + 1:
		fmt.Println("Tomorrow.")
	case tody + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	return
}

func switch_3() {
	// 条件のないswitch
	t := time.Now()
	// switch ture {} と同義,caseで条件分岐(if else ifの代替)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	return
}

func switchtuto() {
	switch_1()
	switch_2()
	switch_3()
	return
}
/// switch end }}}

func main() {
	loop()
	iftuto()
	switchtuto()
}

// NOTE:関数で困ったらGodocする

package main

import (
	"fmt"
	"math"
	"time"
)


// 流用
type myError struct {
	When time.Time
	What string
}
func (err *myError) Error() string {
	return fmt.Sprintf("\nERROR:%T\ntime:%v\nmsg:%s\n", err, err.When, err.What)
}

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
	// float64とかError()の無い型にキャストしないと再帰呼出しの無限ループになるらしい
	// fmtパッケージではError()があればそれを呼び出す様になっている
	// そのため無限再帰でメモリがあふれるみたい
}

func Sqrt(x float64) (float64, error) {
	// DONE:負の数を引数にとった時errorを返すようにする

	// 異常系 流用
//	if x < 0 {
//t.S		return x, &myError{ time.Now(), "x is a negative number" }
//	}

	// 異常系
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := float64(1)
	tmp := x // ループ離脱の比較用
	i := 0 // ループ回数表示のため
	for ; i < 10000; i++ {
		z = z - (z * z - x) / (2 * x)

		if math.Abs(z - tmp) < 0.0000001 { break }

		tmp = z
	}
	fmt.Println("loop cout = ", i)

	// 正常系
	return z, nil
}

// NOTE:
// エラーの書き方
// 異常系の種類ごとに自分でエラー表示を実装して
// ifで返すようにするのかな

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

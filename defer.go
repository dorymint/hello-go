package main

import (
	"fmt"
)


func defer_1() {
	// deferは関数の実行を呼び出し元のreturnまで遅延させる
	// deferに渡した引数はすぐに評価されるらしい
	defer fmt.Println("use defer1")// 5

	{
		defer fmt.Println("scope1") // 4 
		fmt.Println("scope2") // 1
	}

	defer fmt.Println("use defer2") // 3
	fmt.Println("hello") // 2

	// 呼び出し順を見るとdeferはstackに積んでるっぽい
	return
}

// Stacking defers
func defer_2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	return
}


func main(){

	defer_1()
	defer_2()

}



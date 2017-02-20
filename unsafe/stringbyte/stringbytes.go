package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "hello world!!"
	fmt.Printf("len %v, %T, %s\n", len(str), str, str)

	// ポインタ演算でバイトキャスト
	vec := *(*[]byte)(unsafe.Pointer(&str))
	fmt.Printf("len %v, %T, %s\n\n", len(vec), vec, vec)



	cast := func(n uint){
		const (
			red = "\x1b[31m"
			cyan = "\x1b[36m"
			yellow = "\x1b[33m"
			reset = "\x1b[0m"
		)
		fmt.Printf("drop check, %#x\n", n)
		var root = uint32(n)
		fmt.Printf("root:%s%T, %d%s\n", cyan, root, root, reset)
		fmt.Printf("root:bit %s%032b,%s\n", cyan,  root, cyan)
		fmt.Println(yellow, "CAST", reset)
		var drop = *(*uint8)(unsafe.Pointer(&root))
		fmt.Printf("drop:%s%T, %d%s\n", red,drop,  drop, reset)
		fmt.Printf("drop:bit %s%032b,%s\n\n", red, drop, reset)
	}
	cast(0xffff)
	cast(0xff00)
	cast(0xffffffff)
	cast(0xffff0000)
}

// unsafe はガベ-ジコレクトで補足されない
// ポインタを引き出してスコ-プを抜けると
// もとの変数がガベ-ジコレクトされるかも
// スコ-プ内で参照が残っている間に処理を終わらせるなら大丈夫かも
// unsafe な処理の後にもとの変数を参照するように組めば gc されないはず?

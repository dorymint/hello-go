package main

import (
	"fmt"
)

// 空文字の取り扱い
func stringCack() {
	fmt.Println("stringCack")

	s := string("Go tutorial start")
	fmt.Println(s)

	empty := string("")
	fmt.Printf("%s %T %v %d", empty, empty, empty, empty)
	fmt.Println(empty)

	// 空文字はruneに変換できないerror吐いた
	// rune_empty = rune("")
	// fmt.Printf("%s %T %v %d", empty, empty, empty, empty)
	// fmt.Println(empty)

	var key string
	for _, r := range s {
		if r != ' ' {
			key += string(r)
			key += ""
			fmt.Println(key)
		}
		if string(r) == "" {
			fmt.Println("empty!")
		}
	}
	return
}

// 美しすぎる
func love() string { return string('生' & '死') }

func typeCeck() {
	var i byte = 'A'
	fmt.Printf("%v, %T\n", i, i)
	i++
	fmt.Printf("%v, %T\n", i, i)

}

// goの動作でちょっと気になったことの確認用
func main() {

	typeCeck()
	stringCack()
	fmt.Println(love())
}

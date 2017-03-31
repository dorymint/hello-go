package main

import (
	"fmt"
	"reflect"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

// 空文字の取り扱い
func stringCheck() {
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

// love
func love() string { return string('生' & '死') }

func typeCheck() {
	var i byte = 'A'
	fmt.Printf("%v, %T\n", i, i)
	i++
	fmt.Printf("%v, %T\n", i, i)
}

// 除算テスト
func quotientRemainder() {
	fmt.Println("quotientRemainder")
	result := float64(0 % 3)
	fmt.Println(result)
}

// slice
func slice() {
	fmt.Printf("%q\n", []string{})
	ss := []string{"h","e","l","l","o"}
	fmt.Printf("ss    , %q\n", ss)
	fmt.Printf("ss[0:], %q\n", ss[0:])
	fmt.Printf("ss[1:], %q\n", ss[1:])
	fmt.Printf("ss[5:], %q\n", ss[5:])
	if ss[5:] == nil {
		fmt.Println("hi")
	}
	fmt.Println("TypeOf(ss[5:]) = ", reflect.TypeOf(ss[5:]))
	fmt.Println("len(ss[5:]) ,", len(ss[5:]))
	//fmt.Printf("%q\n", ss[6:]) // panic, invalid
	fmt.Printf("ss    , %q\n", ss)
	fmt.Printf("ss[:0], %q\n", ss[:0])
	fmt.Printf("ss[:1], %q\n", ss[:1])
	fmt.Printf("ss[:5], %q\n", ss[:5])
}

func main() {
	split("typeCheck")
	typeCheck()
	split("stringCheck")
	stringCheck()
	split("love")
	fmt.Println(love())
	split("quotientRemainder")
	quotientRemainder()
	split("slice")
	slice()
}

package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

func map_1() {
	fmt.Println("map_1")

	// data
	type Vertex struct {
		Lat, Long float64
	}

	// map declare
	var m map[string]Vertex

	// map make
	m = make(map[string]Vertex)

	// append literal data
	m["Bell Labs"] = Vertex{
		40.68443, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// map literal...{ key :(コロン) {要素}, }
	fmt.Println("\nmap literal")
	var m_2 = map[string]Vertex{
		"Ruby": Vertex{
		40.68443, -74.39967,
		},
		"Google": Vertex{
			37.42202, 122.08408,
		},
	}
	fmt.Println(m_2)

	// mapに渡すトップレベルの型が単純な型名なら
	// literalの要素から方が推定できるため型名を省略できる
	fmt.Println("\n型名の省略")
	var m_3 = map[string]Vertex{
		// Vertexを省略
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m_3)

	return
}

// mutating
func map_2() {
	fmt.Println("map_2")

	m := make(map[string]int)

	// 追加
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// 変更
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 存在確認...戻り値を2つ持つ、keyが存在しなければゼロ値と不存在を表すfalseが返る
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	return
}

// golang.org/x/tour/wc からテストを呼んで実装した関数を試す
func wordCount() {

	// param: string
	// 引数stringに出てくる各単語の出現回数のマップを返す
	wordCountMain := func (s string) map[string]int {
		// stringを文字で扱う時はruneを使う
		// stringのrangeを取るとruneが返る

		result := make(map[string]int)
		var key string

		for _, char := range(s) {
			if char != ' ' {
				key += string(char)
			} else {
				result[key]++
				key = ""
			}
		}
		// ループで空文字をトリガーにした処理が最後に実行されないため取り敢えず
		if key != "" { result[key]++ }

		return result
	}

	wc.Test(wordCountMain)

	fmt.Println()
	wordCountMain("hello world")
	return
}

func main() {

	map_1()
	fmt.Println()
	map_2()
	fmt.Println()

	wordCount()
	fmt.Println()


}


// 2つの二分木が同一の順列で数列を保持しているか確認する

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)
// golang.org/x/tour/tree
// NOTE:use :Godoc golang.org/x/tour/tree
// type Tree struct
// tree.New(k) で値を持つランダムに構造化した二分木を作れる

// DONE:Walk test
// go Walk(tree.New(1), ch)
// TODO:Samu test
// Samu(tree.New(1), tree.New(1)) true
// Samu(tree.New(1), tree.New(2)) false


// Walk tree sending all values
// from tree to channel ch.
// DONE:implement
// in oder
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree, ch chan int)
	// implement
	walk = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			walk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right, ch)
		}
	}

	// work
	// TODO:callstackが溜まりまくりそう、もっといい方法を考えたい
	walk(t, ch)
	close(ch)
}

// Same determines whether the tree
// TODO:implement
func Same(t1, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// comparison test
	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2

		// value check
		if x != y { return false }

		// length check
		if (ok1 || ok2) && !(ok1 && ok2) { return false }

		// pass a test
		if !ok1 && !ok2 { return true }
	}
}


func main() {

	// Walk test
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for x := range ch {
		fmt.Println(x)
	}

	// Same test
	result := Same(tree.New(1), tree.New(1))
	fmt.Println("Same result = ", result)

	result2 := Same(tree.New(2), tree.New(1))
	fmt.Println("Same result2 = ", result2)

}

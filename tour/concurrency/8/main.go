// Exercise: Equivalent Binary Trees.
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var f func(*tree.Tree, chan int)
	f = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		f(t.Left, ch)
		ch <- t.Value
		f(t.Right, ch)
	}
	f(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		i1, ok1 := <-ch1
		i2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 {
			return false
		}
		if i1 != i2 {
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Printf("want true,  out: %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("want false, out: %v\n", Same(tree.New(1), tree.New(2)))
}

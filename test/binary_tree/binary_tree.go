package main

import (
	"fmt"
	//"golang.org/x/tour/tree"
	"math/rand"
	"os"
	"time"
)

/*

// 葉に同じ順序の値を保持する異なる二分木がある
// 2つの二分木が同じ順序を保持しているか?
// goroutineを使いつつ確認する

// 二分木構造
type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

// 二分木で保存する数列 fib
// { 1, 1, 2, 3, 5, 8, 13 }

// preoder, postoder, inoder
// 検索はinoderを使う

*/

// Tree 二分木 親への参照は付けてない単方向
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree

	// NOTE:データを足す時はNewTreeを変更するだけで良くする
}
// TODO:structにするかもう少し考えたい
var (
	treeSize int
	treeNest int
)
// TODO:これも他に方法を考えたい
func showInit() {
	treeSize = 0
	treeNest = -1
}

// NewTree option string are "new" or "L" or "R". select for insert node point
func NewTree(t *Tree, n int, s string) *Tree {
	// NOTE:Treeのデータを足す時はこちらを変更
	if s == "L" && t != nil {
		return &Tree{t.Left, n, nil}
	}
	if s == "R" && t != nil {
		return &Tree{nil, n, t.Right}
	}
	if s == "new" {
		// TODO:t.preNode追加した時用の分岐がある
		if t == nil {
			return &Tree{nil, n, nil} // t.preNode &Tree{nil, n, nil, nil}
		}
		return &Tree{nil, n, nil} // t.preNode &Tree{nil, n, nil, t.preNode}
	}
	fmt.Fprintf(os.Stderr, "NewTree \"%s\" is unknown option\n", s)
	fmt.Fprintf(os.Stderr, "if you use \"R\" or \"L\" are need pointer %T\n", t)
	return nil
}

// InsertL 左にnode挿入
func (t *Tree) InsertL(n int) *Tree {
	return t.insertLR(n, "L")
}

// InsertR 右にnode挿入
func (t *Tree) InsertR(n int) *Tree {
	return t.insertLR(n, "R")
}

// 変なoption投げない様しとく
func (t *Tree) insertLR(n int, s string) *Tree {
	switch s {
	case "L":
		if t.Left != nil {
			t.Left = NewTree(t, n, "L")
			return t.Left
		}
		t.Left = NewTree(t, n, "new")
		return t.Left
	case "R":
		if t.Right != nil {
			t.Right = NewTree(t, n, "R")
			return t.Right
		}
		t.Right = NewTree(t, n, "new")
		return t.Right
	default:
		fmt.Fprintf(os.Stderr, "insertLR invalid option %s\n", s)
		os.Exit(1)
		return nil // os.Exit(1)してるけどreturnが無いとエラーになる
	}
}

// Insert 指定したrootから値を比較し昇順左詰めでnodeを作る
// from:"golang.org/x/tour/tree"
// トップダウンで値を入れてる
func Insert(t *Tree, n int) *Tree {
	if t == nil {
		return NewTree(nil, n, "new")
	}
	if n < t.Value {
		t.Left = Insert(t.Left, n)
	} else {
		t.Right = Insert(t.Right, n)
	}
	return t
}

// Show preoder Left start
// treeNest, treeSize はグローバルな変数
// 依存してるのでroot.Show()を呼ぶ前にshowInit()を呼ぶ
// TODO:検索と表示が混ざってる、数値だけ返すようにして表示は分けるべき
func (t *Tree) Show() {
	// TODO:callstack
	defer func(){treeNest--}()
	treeNest++
	treeSize++

	// TODO:research 末尾再帰最適化
	fmt.Printf("%d  nest%d size%d\n", t.Value, treeNest, treeSize)
	if t.Left == nil && t.Right == nil {
		fmt.Println("is endleaf")
		return
	}

	// TODO:callstackが...
	if t.Left != nil {
		fmt.Printf("L ")
		t.Left.Show()
	}
	if t.Right != nil {
		fmt.Printf("R ")
		t.Right.Show()
	}
}

// TODO:test
// test data
var (
	fib = [...]int{0, 1, 1, 2, 3, 5, 8, 13}
)

func addFibLeft(t *Tree) {
	tmp := t
	for _, x := range fib {
		tmp = tmp.InsertL(x)
	}
}
func addFibRight(t *Tree) {
	tmp := t
	for _, x := range fib {
		tmp = tmp.InsertR(x)
	}
}
func addFibRandom(t *Tree) {
	rand.Seed(time.Now().UnixNano())
	tmp := t
	for _, x := range fib {
		if rand.Int()%2 == 0 {
			tmp = tmp.InsertL(x)
		} else {
			tmp = tmp.InsertR(x)
		}
	}
}
func nodeTest() {
	/* root1 */
	root := new(Tree)

	fmt.Println("\n\nroot.Show() ...1")
	addFibLeft(root)
	showInit()
	root.Show()

	fmt.Println("\n\nroot.Show() ...2")
	addFibRight(root)
	showInit()
	root.Show()

	fmt.Println("\n\nroot.Show() ...3")
	addFibRight(root.Left)
	showInit()
	root.Show()

	fmt.Println("\n\nroot.Show() ...4")
	addFibRandom(root.Right)
	showInit()
	root.Show()

	/* root1 end */

	/* root2 end */
	root2 := new(Tree)
	addFibLeft(root2)
	addFibRandom(root2)
	addFibRandom(root2)
	addFibRandom(root2)
	addFibRandom(root2)
	addFibRandom(root2)

	fmt.Println("\n\nroot2.Show()")
	showInit()
	root2.Show()
	/* root2 end */

}

func main() {
	nodeTest()

	fmt.Println("main end")

}


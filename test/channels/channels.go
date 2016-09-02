package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	defer fmt.Println("completed for sum")
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func lazySum(s []int, c chan int) {
	defer fmt.Println("completed for lazySum")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("lazySum called sum()")
	sum(s, c)
}

// channelはバッファに使える
func chanBuf() {
	ch := make(chan int, 2) // buffer capacityを第二引数で指定
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	// fmt.Println(<-ch) // バッファの終端を超えて読むと実行時エラー?

	ch <- 3
	ch <- 4
	fmt.Println(<-ch) // 3
	ch <- 5
	fmt.Println(<-ch) // 4
	fmt.Println(<-ch) // 5
	fmt.Println("len(ch) = ", len(ch)) // lne(ch) buffer内の要素数...可変
	fmt.Println("cap(ch) = ", cap(ch)) // cap(ch) bufferの総容量...固定? 後から拡張できるのか?

	fmt.Println("push 2 elm")
	ch <- 6
	ch <- 7
	// ch <- 8 // buffer length over 実行時エラー
	fmt.Println("len(ch) = ", len(ch))
	fmt.Println("cap(ch) = ", cap(ch))

	// ch が queue になってる、簡単でいい
}

// channelとclose
func channelClose() {
	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		// 送り手が送信する値がなくなったことを示すためにchannelをcloseできる
		close(c)
		// c <- 1 // closeしたchannelへ送信すると実行時エラー
	}


	c := make(chan int, 10)
	c <- 1
	fmt.Printf("channel <c> type = %T\nvar = %v\n", c, c)
	<-c // これでnullにpopできる


	// そのままチャンネルを渡してるけどストリームへの参照と思って良さそう
	go fibonacci(cap(c), c)
	// range c で値とboolが返る、close(c)でfalse
	for i := range c {
		fmt.Println(i)
	}

	// channelのcloseチェック
	v, ok := <-c
	fmt.Println(v, ok) // 閉じてればfalse
	v = <-c
	fmt.Println(v)
}

// select
func selectCase() {

	// fibo
	fibonacci := func(c, quite chan int) {
		x, y := 0, 1
		for {
			// selectは複数のgoroutineを受けて実行可能なチャンネルを走らせる
			// 複数実行可能な操作があればランダムに選択される
			// 選択されたgoroutineが終わるまで他のgroutineはブロックされる?
			select {
			case c <- x:
				x, y = y, x+y
			case <- quite:
				fmt.Println("quite")
				return
			}
		}
	}


	c := make(chan int)
	quite := make(chan int)
	fmt.Println("cap(c)",cap(c),"cap(quite)", cap(quite))
	// buffer 0でも問題ないっぽい?
	// make(chan T) make(chan T, n) は違う?
	// NOTE:違いはない、関数の一番下のNOTEで説明

	// ----------------------------
	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("len and cap", len(c), cap(c)) // 0, 0
			fmt.Println(<-c)
			//fmt.Println("len and cap", len(c), cap(c)) // 0, 0
		}
		quite <- 0
	}()
	fibonacci(c, quite)
	// ----------------------------
	// 2つの関数はgoroutineによって非同期に実行されselectで足を揃える
	//
	// func():
	// fibonacci()のselectでfunc()に送るデータが詰められるのを待つ
	// 待つ場所はfmt.Println(<-c)の <-c でデータを待つことになる
	// データを吐き終えたらquiteに1つデータを入れる
	// <-quite が実行可能になる
	//
	// fibonacci():
	// selectで実行可能なチャンネルを実行する
	// 始め <-quite はデータ待ちで実行できない
	// c <- x が実行される
	// c にデータが入って吐き出し待ちになる
	// func()で <-c されるとselectされる
	// 最後はquiteのデータを吐いてreturnする

	fmt.Println("test")
	test := make(chan int)
	go func(){test <- 10}()
	i, ok := <-test
	fmt.Printf("%T,%v",i, ok)

	// error吐く
	//v, ok := <-c
	//fmt.Println(v, ok)

	// make(chan T)とmake(chan T, n)は単にバッファの容量が違うだけじゃないっぽい?
	//
	// NOTE:理解した、make(chan T) と make(chan T, n)は単に容量が異なるだけである
	// make(chan int) capはゼロ
	// buffer 0なので入力は詰まって処理が止まる
	// 出力側が取り出そうとした時に初めて入力側のbuffer 0のチャンネルを通れる
	// buffer 0のチャンネルを使う時は並行して出力も動かすか
	// あるいは入力を別スレッドにしないと deadlockする...deadlockする
}


// goroutine と channel
func main() {
	s := []int{ 7, 2, 8, -9, 4, 0 }

	fmt.Println("\nsync test")

	// int型のチャンネル
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	go lazySum(s[:], c)


	// <- channel operator
	x, y := <-c, <-c

	fmt.Println("waiting lazySum")
	lazy := <-c
	// 上のチャンネルが同期するまで次に行かない

	fmt.Println(x, y)
	fmt.Println(lazy)

	fmt.Println("\n call chenBuf")
	chanBuf()

	fmt.Println("\n call channelClose")
	channelClose()

	fmt.Println("\n call selectCase")
	selectCase()
}

// NOTE:channel operator は goroutine を同期する

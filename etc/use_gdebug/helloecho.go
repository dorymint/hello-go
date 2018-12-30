// デバッグシンボル付きのコンパイル
// go build -gcflags "-N -l" <dst src>

// gdb <exe file name>
// run, list <n>, break <n>, delete <n>, print <v>,
// info breakpoints...
package main

import (
	"fmt"
	"time"
)


func whileEcho(s string, n int) {
	for i := 0; i < n; i++  {
		fmt.Println(s)
	}
}

// 引数の chan<- int と chan int の違い
// 送信専用: chan<- int
// 受信専用: <-chan int
// chan<- で送信 <-chan で受信を明示してる、単にchanだとどちらも使える
func counting(c chan<- int) {
	fmt.Printf("%p,%p,\n", c, &c)
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
		fmt.Println("hello world")
	}
	close(c)
}


func main() {
	// gdb test1
	fmt.Println("hello")
	whileEcho("world", 10)

	// gdb test2
	msg := "String main"
	fmt.Println(msg)
	bus := make(chan int)

	// gdb で変数の中身を確認するため変更
	msg = "starting a gofunc"

	go counting(bus)
	fmt.Printf("main chan %p, %p\n", bus, &bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	defer func() { fmt.Printf("\n%s\t%s\n", "done!!!", s) }()
	for i := 0; i < 5; i++ {
		// 動作の一時停止 time package で使える
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	// すごい面白い
	// 実行するたびに順番が変わるのは関数が非同期で実行されているから
	go say("world...sub") // サブスレッド
	go say("hello...sub") // サブスレッド
	say("gopher...main") // mainスレッドで実行される

	// mainが終わるとサブスレッドで作業が残っていても一緒に死んでしまう

}

// NOTE:goroutine
// goランタイムで管理されるスレッド
// go をキーワードに新しく作られる
// goの対象は関数オブジェクトだけ?

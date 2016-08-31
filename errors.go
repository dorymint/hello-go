package main

// 少し多めにコメント書きながら確認する

/* stringer と同じように error にも組み込みインターフェースがある */
// type error interface { Error() string }
// fmtが変数を文字列で出力する時errorインターフェースか確認する

import (
	"fmt"
	"time"
)

// errorの形式を決める
type myError struct {
	When time.Time // 時刻を記録できる変数、error発生時の時刻を記録
	What string // errorメッセージ
}
// errorメッセージを作って返す
func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// errorが発生する可能性のある関数とする
func run() error {
	// error処理、実体を作りポインタで返す
	return &myError{
		time.Now(), // 現在時刻の取得
		"it didn't work.",
	}
}

func main() {
	// run()を実行しerrorが返っていれば処理
	// errorのチェックは != nil で行う、変数がnilならエラーは返っていない
	if err := run(); err != nil {
		// 型の確認
		fmt.Printf("%T\n\n", err)
		// メッセージの確認
		fmt.Println(err)
		// 今のところerror型とstring型の違いが見えないけど
		// 多分出力先が >&2 になってそう
	}
}


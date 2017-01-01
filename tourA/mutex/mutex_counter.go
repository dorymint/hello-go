
package main
// リソースへの同時アクセスを制限しなければならない場合


import (
	"fmt"
	"os"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	time.Sleep(time.Second)
	// Lock so only one gouroutine at a time can access the map c.v
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
	// 書き込み中のマップに対して新しい書き込みを制限する
}

// Value return the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map c.v
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
	// NOTE:
	// IncのLockと同じくデータへの同時アクセスを制限している
	// 上のLockが無ければ並列実行のタイミング次第でエラーになる
	// 正常終了してしまう可能性もある
}

func main() {
	const targetNumber int = 10000 // あまり大きくし過ぎるとgoroutineが増えすぎてstackが溢れる
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < targetNumber; i++ {
		go c.Inc("somekey")
	}

	fmt.Print("\x1b[33m")
	for x := c.Value("somekey"); x < targetNumber; x = c.Value("somekey") {
		time.Sleep(time.Microsecond)
		fmt.Fprintf(os.Stderr, "%5v\r", x)
	}
	fmt.Print("\x1b[0m\n")

	fmt.Printf("%s%v%s\n", "\x1b[31m", c.Value("somekey"), "\x1b[0m" ) // 1000
	fmt.Printf("%send of work,%s %v\n", "\x1b[32m", "\x1b[0m",  c.Value("somekey"))
}

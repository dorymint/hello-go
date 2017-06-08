package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	var s string
	buf := bytes.NewBufferString(s)
	rwmutex := new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rwmutex.Lock()
			defer rwmutex.Unlock()
			fmt.Fprintln(buf, "hello:", i)
		}(i)
	}

	wg.Wait()
	fmt.Println(buf)
}

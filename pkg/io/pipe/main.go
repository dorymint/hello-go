// pipe.
package main

import (
	"fmt"
	"io"
	"sync"
	"time"
)

func main() {
	pr, pw := io.Pipe()
	var wg sync.WaitGroup

	// write
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pw.Close()
		for _, m := range []byte("123456") {
			pw.Write([]byte{m})
			time.Sleep(time.Second)
		}
	}()

	// read
	wg.Add(1)
	go func() {
		defer wg.Done()
		buf := make([]byte, 3)
		var err error = nil
		for {
			_, err = io.ReadFull(pr, buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(string(buf))
		}
	}()

	wg.Wait()
}

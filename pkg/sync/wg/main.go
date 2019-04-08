// sync wait group.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go wg.Done()

	wg.Wait()
	wg.Wait()

	fmt.Println("exit")
}

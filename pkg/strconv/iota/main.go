// iota.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	us := []uint{50, 2, 1, 9, 1<<64 - 1}
	fmt.Println("uint slice:", us)
	for _, u := range us {
		fmt.Printf("%s %s\n",
			strconv.FormatUint(uint64(u), 10),
			strconv.Itoa(int(u)))
	}
}

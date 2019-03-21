// format uint.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	uis := []uint64{0, 10, 100, 1000, 10000}
	uis = append(uis, 1, 12, 123, 1234, 12345)
	for i, ui := range uis {
		fmt.Printf("i=%d ui=%d %s\n", i, ui, strconv.FormatUint(ui, 10))
	}
}

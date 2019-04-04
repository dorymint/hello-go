// sync map.
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := new(sync.Map)

	storeAndPrint := func(key interface{}, val interface{}) {
		m.Store(key, val)
		val, ok := m.Load(key)
		fmt.Printf("key:%+v\n", key)
		fmt.Printf("\tval:%+v\n", val)
		fmt.Printf("\tok:%v\n\n", ok)
	}
	storeAndPrint("hello world", 100)
	storeAndPrint(10, []byte("foo bar"))
}

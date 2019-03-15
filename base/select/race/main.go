// select.
package main

import (
	"fmt"
)

func main() {
	chs := make([]chan string, 10)
	for i := range chs {
		chs[i] = make(chan string)
	}

	for i, ch := range chs {
		go func(ch chan string, i int) {
			ch <- fmt.Sprintf("winner is ch[%d]!", i)
		}(ch, i)
	}

	var s string
	select {
	case s = <-chs[0]:
	case s = <-chs[1]:
	case s = <-chs[2]:
	case s = <-chs[3]:
	case s = <-chs[4]:
	case s = <-chs[5]:
	case s = <-chs[6]:
	case s = <-chs[7]:
	case s = <-chs[8]:
	case s = <-chs[9]:
	}
	fmt.Println(s)
}

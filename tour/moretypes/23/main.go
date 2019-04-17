// Exercise: Maps.
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	ss := strings.Fields(s)
	for i := range ss {
		m[ss[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

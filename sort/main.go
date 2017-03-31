package main

import (
	"fmt"
	"sort"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func sortSlice() {
	ul := []uint{50, 2, 1, 9}
	fmt.Println(ul)
	sort.Slice(ul, func(i, j int) bool {
		return ul[i] < ul[j]
	})
	fmt.Println(ul)
}

func main() {
	split("sortSlice")
	sortSlice()
}

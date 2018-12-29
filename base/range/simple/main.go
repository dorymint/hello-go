package main

import (
	"fmt"
)

func main() {
	for i := range make([]struct{}, 10) {
		fmt.Println(i)
	}
}

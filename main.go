// scriptencoding utf-8

package main

import (
	"time"
	"fmt"
)

func main() {

	fmt.Printf("hello go\n")

	str := "hello"
	fmt.Printf("%s\n", str)


	fmt.Println("welcome to the go tuto!")
	fmt.Println("The time is" , time.Now())
	fmt.Println("go run test")
}




// :nnoremap <c-@>gorun :!go run main.go > gorunlog.txt<nl>

package main

import (
	"fmt"
	"path/filepath"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func q1() {
	fmt.Println(filepath.Abs("./t//dfj"))
	fmt.Println(filepath.Abs("./t//dfj/"))
	fmt.Println(filepath.Dir("./t//tdfklj/"))
	fmt.Println(string(filepath.Separator))
}

func main() {
	split("q1")
	q1()
}

package main

import (
	"fmt"
	"path/filepath"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func q1() {
	fmt.Println("separator:", string(filepath.Separator))
	fmt.Println("dir:", filepath.Dir("./t//tdfklj/"))
	fmt.Println(filepath.Abs("./t//dfj"))
	fmt.Println(filepath.Abs("./t//dfj/"))
	fmt.Println(filepath.Abs(""))
}

func main() {
	split("q1")
	q1()
}

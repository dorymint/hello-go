package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func abs() {
	q1 := func(f string) {
		path, err := filepath.Abs(f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("case:", f+"\n", path)
	}
	q1("./")
	q1("./../")
	q1("") // pwd
	q1("/invalid")
	q1("invalid/invalid") // joined current working directory
}

func q1() {
	fmt.Println("separator:", string(filepath.Separator))
	fmt.Println("dir:", filepath.Dir("./t//tdfklj/"))
	fmt.Println(filepath.Abs("./t//dfj"))
	fmt.Println(filepath.Abs("./t//dfj/"))
	fmt.Println(filepath.Abs(""))
}

func main() {
	log.SetFlags(log.Lshortfile)
	split("abs")
	abs()
	split("q1")
	q1()
}

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

func main() {
	log.SetFlags(log.Lshortfile)
	split("abs")
	abs()
}

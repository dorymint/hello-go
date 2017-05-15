package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	write := func(fname string) {
		if err := ioutil.WriteFile(fname, []byte("hello"), 0600); err != nil {
			log.Println(err)
		}
	}
	if err := os.Chdir("t/nest"); err != nil {
		log.Fatal(err)
	}
	write("..")
	write("hello.txt")
	write("./hello1.txt")
	write("../hello2.txt")
}

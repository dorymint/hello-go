package main

import (
	"log"
	"os"
)

func main() {
	err := os.Mkdir("./modedir", os.ModeDir)
	if err != nil {
		log.Println(err)
	}
	err = os.Mkdir("./withperm", os.ModeDir|0666)
	if err != nil {
		log.Println(err)
	}
	err = os.Mkdir("./perm", 0777)
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

var dir = flag.String("d", "", "")

func main() {
	flag.Parse()
	fmt.Printf("dir: %q\n", *dir)

	fpath, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fpath:", fpath)

	fpath, err = filepath.Abs(fpath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fpath:", fpath)

	userinfo, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userinfo)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pwd:", pwd)

	target := []string{
		"target.txt",
		"../target.txt",
		"../target.txt/../",
		"../target.txt/../tset.txt",
		"../target.txt/..//tset.txt//",
	}
	split("data")
	for _, x := range target {
		fmt.Println(x)
	}
	split("Join")
	for _, x := range target {
		fmt.Println("      :", filepath.Join(pwd, x))
	}
	split("base")
	for _, x := range target {
		fmt.Println("      :", filepath.Base(x))
	}
}

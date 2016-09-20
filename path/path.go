package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// str is from function infomation.
func fatalIF(str string, err error) {
	if err != nil {
		log.Fatalf("%s:%v\n", str, err)
	}
}

// flags
var (
	root = flag.String("root", "./", "Specify search root")
)

func init() {
	var err error
	flag.Parse()
	*root, err = filepath.Abs(*root)
	fatalIF("init", err)
}

func proccess() error {
	return nil
}

func main() {
	fmt.Println(*root)

	// walk func test
	dir, err := os.Open(*root)
	fatalIF("open *root:", err)
	defer dir.Close()

	dirname, err := dir.Readdir(0)
	fatalIF("readdir:", err)



	fmt.Fprint(os.Stdin, "echo hello")


	for _, x := range dirname {
		if x.IsDir() {
			fmt.Println(x.Name())
		}
	}

	// filepath.Dir is get current directory
	fmt.Println(filepath.Dir(*root))
}

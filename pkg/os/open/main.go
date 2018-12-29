package main

import (
	"errors"
	"fmt"
	"os"
)

func pwd() (dir string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	return
}

func snipWorkenv() {
	var err error
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("HOST:", host)

	f, err := os.Open(pwd())
	if err != nil {
		panic(err)
	}

	info, err := f.Readdir(0)
	if err != nil {
		panic(err)
	}
	for i, x := range info {
		fmt.Println("fileinfo", i, x.Name(), x.Size(), x.Mode())
	}

	pwd()

	// tempdir test
	dirname := os.TempDir()
	dir, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := dir.Close()
		if err != nil {
			panic(err)
		}
	}()
	fmt.Println(dir.Readdir(1))

	err = os.Chdir(dirname)
	if err != nil {
		panic(err)
	}

	pwd()

	fmt.Println("GOROOT", os.Getenv("GOROOT"))

	err = os.Chdir("/")
	fmt.Println(err)
	pwd()
}

func snipFileRead() error {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir, err := os.Open(root)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	// reaadはファイル内のreadポインタを動かしたままにする
	fmt.Println("info to dir")
	dirnames, err := dir.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	for i, x := range dirnames {
		fmt.Println(i, x)
	}

	// readポインタが動いているのでdirenamesが見つけられない
	fmt.Println("info to verbose")
	info, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	for i, x := range info {
		fmt.Println(i, x.IsDir(), x.Name())
	}

	return errors.New("test")
}

func main() {
	snipWorkenv()

	snipFileRead()
}

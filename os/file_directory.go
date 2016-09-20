package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func init() {

}

func fatalIF(str string, err error) {
	if err != nil { log.Fatalf("%s:%q\n", str, err) }
}

func pwd() (dir string) {
	dir, err := os.Getwd()
	fatalIF("pwd",err)
	fmt.Println(dir)
	return
}

func snipWorkenv() {
	var err error
	host, err := os.Hostname()
	fatalIF("Hostname", err)
	fmt.Println("HOST:", host)

	f, err := os.Open(pwd())
	fatalIF("open+pwd()", err)

	info, err := f.Readdir(0)
	fatalIF("Readdir", err)
	for i, x := range info {
		fmt.Println("fileinfo", i, x.Name(), x.Size(), x.Mode())
	}

	pwd()

	// tempdir test
	dirname := os.TempDir()
	dir, err := os.Open(dirname)
	fatalIF("open"+dirname, err)
	defer func() { errclose := dir.Close(); fatalIF("close", errclose) }()
	fmt.Println(dir.Readdir(1))

	err = os.Chdir(dirname)
	fatalIF("cd", err)

	pwd()

	fmt.Println("GOROOT", os.Getenv("GOROOT"))

	err = os.Chdir("/")
	fmt.Println(err)
	pwd()
}

func snipFileRead() error {
	root, err := os.Getwd()
	fatalIF("os.Getwd", err)
	dir, err := os.Open(root)
	fatalIF("os.Open", err)
	defer func() { fatalIF("TODOGather", dir.Close()) }()

	// reaadはファイル内のreadポインタを動かしたままにする
	fmt.Println("info to dir")
	dirnames, err := dir.Readdirnames(0)
	fatalIF("Readdirnames", err)
	for i, x := range dirnames {
		fmt.Println(i, x)
	}

	// readポインタが動いているのでdirenamesが見つけられない
	fmt.Println("info to verbose")
	info, err := dir.Readdir(0)
	fatalIF("f.Readdir", err)
	for i, x := range info {
		fmt.Println(i, x.IsDir(), x.Name())
	}

	return errors.New("test")
}

func main() {
	snipWorkenv()

	snipFileRead()
}

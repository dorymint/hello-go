package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read(w io.Writer, path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(b))
	return nil
}

func deleteContent(f *os.File) {
	if err := f.Truncate(0); err != nil {
		panic(err)
	}
}

func main() {
	const first = "hello "
	const second = "world "

	f, err := ioutil.TempFile("", "golang_write_test")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	var s string
	buf := bytes.NewBufferString(s)
	writeRead := func(ws string) {
		if _, err := f.WriteString(ws); err != nil {
			panic(err)
		}
		read(buf, f.Name())
	}
	writeRead(first)
	writeRead(second)
	fmt.Println(buf)// hello hello world

	buf.Reset()
	deleteContent(f)

	writeRead(first)
	deleteContent(f)
	writeRead(second)
	fmt.Println(buf) // hello world
}

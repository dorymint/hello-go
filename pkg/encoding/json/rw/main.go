// write to file
// read from file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var DIR = func() string {
	p := filepath.Join("testdata")
	fi, err := os.Lstat(p)
	if err != nil {
		panic(err)
	}
	if !fi.IsDir() {
		panic("is not directory " + p)
	}
	return p
}()

func readJSON() {
	// JSONのkeyが大文字小文字の違いだけならstructにタグを書かなくてもマッピングしてくれる
	type Task struct {
		Name        string
		Description string
		Command     string
	}
	var ts []Task

	p := filepath.Join(DIR, "read.json")
	b, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &ts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ts)
}

func writeJSON() {
	// 書き込むときはタグを付けないとそのまま大文字で出力されてしまう
	type Task struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Command     string `json:"command"`
	}
	var ts []Task

	ts = append(ts, Task{
		Name:        "write",
		Description: "short description",
		Command:     "write",
	})

	var (
		out []byte
		err error
	)

	// marshal
	out, err = json.Marshal(ts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", out)

	// with indent
	out, err = json.MarshalIndent(ts, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", out)

	fmt.Println(`
Pick for write
  ioutil.WriteFile("path/file", out, 0600)
  fmt.Fprint(*os.File, out)
  *os.File.Write(out)
  etc...`)
}

func main() {
	fmt.Println("readJSON:")
	readJSON()
	fmt.Println()

	fmt.Println("writeJSON:")
	writeJSON()
	fmt.Println()
}

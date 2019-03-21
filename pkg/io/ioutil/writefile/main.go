package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	tmpf, err := ioutil.TempFile("", "ioutil")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.Remove(tmpf.Name()); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Printf("%s is removed\n", tmpf.Name())
	}()

	content := fmt.Sprintf("filepath:%q\nwrited by ioutil.WriteFile\n", tmpf.Name())
	if err := ioutil.WriteFile(tmpf.Name(), []byte(content), 0600); err != nil {
		if err != nil {
			panic(err)
		}
	}

	if _, err := io.Copy(os.Stdout, tmpf); err != nil {
		panic(err)
	}
}

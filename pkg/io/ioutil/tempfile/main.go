// temp file.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	tmpf, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		panic(err)
	}
	// tmpf.Name is provided fullpath
	defer os.Remove(tmpf.Name())

	fmt.Printf("%#v\n\n", tmpf)
	fmt.Printf("tmpf.Name():%q\n", tmpf.Name())
	fmt.Printf("filepath.IsAbs:%v\n", filepath.IsAbs(tmpf.Name()))
}

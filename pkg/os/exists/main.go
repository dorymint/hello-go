// os.IsExists, os.IsNotExists.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, err := ioutil.TempDir("", "testexists")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	f, err := ioutil.TempFile(dir, "")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// always false if err is nil
	fmt.Println("os.Stat(existFile)")
	_, err = os.Stat(f.Name())
	// file exists but return false. because err is nil
	fmt.Printf("\tIsExist: %v\n", os.IsExist(err))         // false
	fmt.Printf("\tIsNotExist: %v\n\n", os.IsNotExist(err)) // false

	// if only check the exists
	_, err = os.Stat(f.Name())
	if err == nil {
		fmt.Printf("%s is exist\n\n", f.Name())
	}

	// specified O_EXCL is for return error if file exist
	fmt.Println("os.OpenFile(existFIle) with O_CREATE|O_EXCL")
	ff, err := os.OpenFile(f.Name(), os.O_CREATE|os.O_EXCL, 0600)
	if err == nil {
		ff.Close()
		panic("want error but nil")
	}
	fmt.Printf("\tIsExist: %v\n", os.IsExist(err))         // true
	fmt.Printf("\tIsNotExist: %v\n\n", os.IsNotExist(err)) // false

	err = os.Remove(f.Name())
	if err != nil {
		panic(err)
	}

	fmt.Println("os.Stat(missingFile)")
	_, err = os.Stat(f.Name())
	fmt.Printf("\tIsExist: %v\n", os.IsExist(err))         // false
	fmt.Printf("\tIsNotExist: %v\n\n", os.IsNotExist(err)) // true
}

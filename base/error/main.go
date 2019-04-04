// error.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, err := ioutil.TempDir("", "golang_error_handling")
	if err != nil {
		panic(err)
	}
	err = os.Remove(dir)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(dir)

	fmt.Printf("%T\n", err)
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("%T\n", e.Err)
	}

	fmt.Println(os.IsPermission(err)) // false
	fmt.Println(os.IsNotExist(err))   // true
}

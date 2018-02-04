package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func scan() {
	var i int
	fmt.Print("scan(): input a integer:>")
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("scan(): i=%d\n", i)
}

func fscan(r io.Reader, w io.Writer) (int, error) {
	var i int
	fmt.Fprint(w, "fscan(): input a integer:>")
	if _, err := fmt.Fscan(r, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func sscan(s string) (int, error) {
	var i int
	if _, err := fmt.Sscan(s, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func main() {
	scan()

	if i, err := fscan(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("fscan(): i=%d\n", i)
	}

	if i, err := sscan("4321"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("sscan(): i=%d\n", i)
	}
}

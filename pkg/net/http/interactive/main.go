package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "rootHandler")
}

func interactive(w io.Writer, r io.Reader) {
	sc := bufio.NewScanner(r)
	fmt.Fprintln(w, "starting repl")
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			panic(err)
		}
		switch txt := sc.Text(); txt {
		case "q", "exit", ":q":
			return
		default:
			fmt.Fprintln(w, "echo:", txt)
		}
	}
	panic("scan failed")
}

func main() {
	http.HandleFunc("/", rootHandler)
	go func() {
		// TODO: is sefe?
		interactive(os.Stdout, os.Stdin)
		os.Exit(1)
	}()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

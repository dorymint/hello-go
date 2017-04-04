package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func read(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		log.Println(sc.Err())
		return ""
	}
	return sc.Text()
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func outNew(out *os.File, prefix string) chan<- string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Fprint(out, prefix)
			fmt.Fprintln(out, <-ch)
		}
	}()
	return ch
}

func interactive(pre string) error {
	ch := outNew(os.Stdout, pre)
	errch := outNew(os.Stderr, "")

	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			return fmt.Errorf("interactiv(): %v", sc.Err())
		}
		switch in := strings.TrimSpace(sc.Text()); in {
		case "exit", "q", ":q", "quit":
			fmt.Println("exit")
			return nil
		case "get":
			if b, err := get(read("rul>")); err != nil {
				errch <- err.Error()
			} else {
				ch <- string(b)
			}
		default:
			ch <- sc.Text()
		}
	}
	return fmt.Errorf("interactiv(): fatal")
}

func main() {
	err := interactive("repl>")
	if err != nil {
		log.Fatal(err)
	}
}

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
	if strings.HasPrefix(url, "http") == false {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func outNew(out *os.File, prefix string) chan<- string {
	ch := make(chan string)
	go func() {
		for {
			_, err := fmt.Fprint(out, prefix)
			if err != nil {
				close(ch)
				return
			}
			_, err = fmt.Fprintln(out, <-ch)
			if err != nil {
				close(ch)
				return
			}
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
			b, err := get(read(pre + ":get:rul>"))
			if err != nil {
				errch <- err.Error()
				continue
			}
			ch <- string(b)
		default:
			ch <- sc.Text()
		}
	}
	return fmt.Errorf("interactiv(): fatal")
}

var s = `
start repl!
commands
[exit] [q] [:q] [quit]	stop the repl
[get]	get web content from url

`
func main() {
	fmt.Println(s)
	err := interactive("repl>")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	gomem "github.com/dory/hello-go/proto/gomem"
)

// Return do not close channel
func outNew(prefix string) chan<- string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Print(prefix)
			fmt.Println(<-ch)
		}
	}()
	return ch
}

// Modify
func modNew(outch chan<- string, prefix string) chan<- string {
	inputch := make(chan string)
	go func() {
		for {
			outch <- fmt.Sprintf("%v:%v", prefix, <-inputch)
		}
	}()
	return inputch
}

func read(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		log.Fatalf("echo(): %v", sc.Err())
	}
	return sc.Text()
}

// {Input > Modify > Output}
func interactive() {
	outch := outNew("gomem>")
	ch := modNew(outch, "hello channel")

	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			log.Fatalf("interactive(): %v", sc.Err())
		}
		// TODO: add commands
		switch in := strings.TrimSpace(sc.Text()); in {
		case "exit", "q", ":q", "quit":
			fmt.Println("exit")
			return
		case "get":
			// TODO: implementation
			//gomem.Get(read(), read())
			//interactive()
			fmt.Println("TODO: implementation")
			outch <- read("first:") + read("second:")
		default:
			ch <- sc.Text()
		}
	}
}

func main() {
	fmt.Println(gomem.AddNum(2, 3))
	interactive()
}

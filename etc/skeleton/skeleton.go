package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Return do not close channel
func displayDaemon() chan<- string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	return ch
}

// {Input > modify > Output}
func modifyDaemon(output chan<- string, prefix string) chan<- string {
	inputch := make(chan string)
	go func() {
		for {
			output <- fmt.Sprintf("%v:%v", prefix, <-inputch)
		}
	}()
	return inputch
}

// Proc!!
func proc() {
	outch := displayDaemon()
	ch := modifyDaemon(outch, "hello channel")

	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			log.Fatalf("proc:%v", sc.Err())
		}
		switch in := sc.Text(); in {
		case "exit":
			return
		case "happy":
			ch <- sc.Text() + " new year!"
		default:
			ch <- sc.Text()
		}
	}
}

func main() {
	proc()
}

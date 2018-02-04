package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	go func() {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			if err := sc.Err(); err != nil {
				log.Fatal(err)
			}
			switch sc.Text() {
			case "exit":
				os.Exit(0)
			default:
				fmt.Println(sc.Text())
			}
		}
		fmt.Println("fail sc scan")
		os.Exit(1)
	}()
	fmt.Println("input accept: time limit of one minute")
	time.Sleep(time.Minute)
}

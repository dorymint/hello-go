// echo loop
package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"./client"
	"./server"
)

func main() {
	port := string(":8080")
	ch := make(chan string)
	go func() {
		if err := server.Listen(port); err != nil {
			log.Fatal("main:server:", err)
		}
	}()
	time.Sleep(time.Second)
	go func() {
		for {
			if err := client.Client(port, ch); err != nil {
				log.Fatal("main:client:", err)
			} else {
				log.Println("CLIENT: call agein")
				continue
			}
		}
	}()

	// main input loop
	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			log.Fatalf("main: %v", sc.Err())
		}
		ch <- sc.Text()
	}
	log.Println("exit main loop")
}


// web client

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)


func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("connect error", err)
	}
	defer conn.Close()

	// simple read
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	// simple write
	conn.Write([]byte("from client: hello world!!"))
	time.Sleep(time.Second)
}

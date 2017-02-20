
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
	fmt.Printf("%s\n%s\n", "--- BUFFER ---", string(buffer))

	// simple write
	conn.Write([]byte("FROM CLIENT: hello world!!"))
	time.Sleep(time.Second)
}

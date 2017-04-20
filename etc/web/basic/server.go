

// web server
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error tcp :8080")
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("ln.Accept error:", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		log.Println(conn.RemoteAddr(), "close")
		conn.Close()
	}()
	log.Println("hello new connection!")
	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.RemoteAddr())

	// simple write
	conn.Write([]byte("from server to client: send to connection!!"))


	// simple read
	buffer := make([]byte, 1024)
	fmt.Println(len(buffer), cap(buffer))
	conn.Read(buffer)
	fmt.Println(string(buffer), "\n")
	time.Sleep(time.Second)
}


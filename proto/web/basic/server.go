
// web server

package main

import (
	"fmt"
	"io"
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
		fmt.Println(conn.LocalAddr())
		fmt.Println(conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println(conn.RemoteAddr(), "is closed\n")
	}()
	log.Println("hello new connection!")
	//io.Copy(conn, conn)

	// simple write
	conn.Write([]byte("from server: send to connection!!"))


	// simple read
	buffer := make([]byte, 1024)
	fmt.Println(len(buffer), cap(buffer))
	conn.Read(buffer)
	fmt.Println(string(buffer), "\n")
	time.Sleep(time.Second)
}


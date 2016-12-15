
// web server

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error tcp :8080")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("ln.Accept error:", err)
		}
		fmt.Println(conn.LocalAddr())
		fmt.Println(conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println(conn.RemoteAddr(), "is closed")
	}()

}


package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// Listen run the top process of server
// use tcp listen
// err=net.Listen
// port=":<No.>"
// example: Listen(":8080")
func Listen(port string) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("SERVER Listen: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		log.Println("SERVER:", conn.LocalAddr(), " | ", conn.RemoteAddr(), "DISCONNECT")
		conn.Close()
	}()
	log.Println("SERVER:", conn.LocalAddr(), " | ", conn.RemoteAddr(), "CONNECTED")
	if err := conn.SetDeadline(time.Now().Add(time.Second * 30)); err != nil {
		log.Println("SERVER:conn.SetDeadline: ", err)
		return
	}
	// TODO: implementation echoback
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		// line には '\n' まで含まれてるっぽい
		line, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("SERVER:", err)
			return
		}

		conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
		fmt.Fprintf(conn, "%s", line)
	}
}

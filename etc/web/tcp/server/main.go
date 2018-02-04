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
	log.Println("START LISTEN: ", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("ln.Accept error:", err)
			continue
		}
		log.Println("conn.RemoteAddr().String()", conn.RemoteAddr().String())
		log.Println("conn.RemoteAddr().Network()", conn.RemoteAddr().Network())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println(conn.LocalAddr(), " | ", conn.RemoteAddr(), " IS CLOSED")
	}()

	fmt.Println("CONNECTION: ", conn.LocalAddr(), " | ", conn.RemoteAddr())
	if err := conn.SetDeadline(time.Now().Add(time.Second*2)); err != nil {
		log.Println("conn.SetDeadline: ", err)
		return
	}

	if _, err := conn.Write([]byte("FROM SERVER: hello world!\n")); err != nil {
		log.Println(err)
		return
	}
	buf := make([]byte, 128)
	time.Sleep(time.Second)
	if _, err := conn.Read(buf); err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s\n%s\n", "--- BUFFER ---", string(buf))
	fmt.Printf("\n%s\n% x\n", "-- BUFFER BYTE ---", buf)
	fmt.Println("echo")
	time.Sleep(time.Second)
	conn.Write([]byte("--- FROM SERVER ECHO ---\n" + string(buf) + "\n"))
	fmt.Println("time waite 5s")
	time.Sleep(time.Second*5)
}

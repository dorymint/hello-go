package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// やたらと実行して迷惑をかけないように
	// conn, err := net.Dial("tcp", "golang.org:80")
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

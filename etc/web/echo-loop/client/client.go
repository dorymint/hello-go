package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// Client main process, use tcp
// example: Client(":8080")
// retrun error = net.Dial || echoback
func Client(port string, ch chan string) error {
	for {
		fmt.Println("CLIENT: wait input from ch :>")
		buf := <-ch
		log.Println("ECOH SRC:", buf)
		switch buf {
		case "exit", "EXIT", "quit", "QUIT", "q", ":q":
			log.Println("CLIENT: exit echoback")
			return fmt.Errorf("exit")
		case "giko":
			fmt.Println("(,,ﾟДﾟ)")
			return nil
		default:
			conn, err := net.Dial("tcp", port)
			if err != nil {
				return err
			}
			if err := echoback(conn, buf); err != nil {
				return err
			}
		}
	}
	return nil
}

// DO SOMETHING
func echoback(conn net.Conn, call string) error {
	defer func() {
		log.Println("CLIENT:", conn.LocalAddr(), " | ", conn.RemoteAddr(), "DISCONNECT")
		conn.Close()
	}()
	log.Println("CLIENT:", conn.LocalAddr(), " | ", conn.RemoteAddr(), "CONNECTED")
	conn.SetDeadline(time.Now().Add(time.Minute * 10))

	// TODO: コネクションを使いまわして server とやり取りできるように
	//     : 入力と処理、表示の分割

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	fmt.Fprintf(conn, "%s\n", call) // '\n' read delim
	log.Println("CLIENT: DEBUG MES: send", call)

	conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	line, err := bufio.NewReader(conn).ReadString('\n')
		// ReadString は見つけた delim も含めて入れるっぽい
	if err != nil {
		log.Println("echoback:", err)
		return err
	}
	fmt.Println("--- CLIENT ---")
	fmt.Printf("result: %s\n% x\n", line, []byte(line))
	// line に改行まで含まれてる
	return nil
}

package main

import (
	"fmt"
	"io"
	"time"
)

// code from
// http://stackoverflow.com/questions/29060922/reading-from-stdin-in-golang
func stream() {
	input, output := io.Pipe()

	go func() {
		defer output.Close()
		for _, m := range []byte("123456") {
			output.Write([]byte{m})
			time.Sleep(time.Second)
		}
	}()

	message := make([]byte, 3)
	_, err := io.ReadFull(input, message)
	for err == nil {
		fmt.Println(string(message))
		_, err = io.ReadFull(input, message)
	}
	if err != io.EOF {
		panic(err)
	}
}

func main() {
	stream()
}

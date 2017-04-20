package main

import (
	"../../server"
	"log"
)

func main() {
	if err := server.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"log"
)

func main() {
	c, err := colorful.Hex("#ffffff")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	fmt.Println(c.Hex())
	fmt.Println(c.RGB255())
}

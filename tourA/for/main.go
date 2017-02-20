package main

import (
	"log"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		x := rand.Int()
		log.Printf("%p %v\n", &x, x)
	}
}

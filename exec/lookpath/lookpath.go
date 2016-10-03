package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	hello, err := exec.LookPath("todogotcha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello)
	fmt.Printf("%T\n", hello)

	hello, err = exec.LookPath("todogotcha!!!")
	if err != nil {
		log.Fatalf("name:%verr:%v", hello, err)
	}
}

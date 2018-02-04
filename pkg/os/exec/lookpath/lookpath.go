package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	lookEXE, err := exec.LookPath("todogotcha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lookEXE)
	fmt.Printf("%T\n", lookEXE)

	// intentional error. invalid file path
	lookEXE, err = exec.LookPath("todogotcha!!!")
	if err != nil {
		fmt.Println("intentional error")
		log.Printf("name:%verr:%v", lookEXE, err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Fprint(os.Stdout, "test")

	time.Sleep(time.Second)
	fmt.Fprint(os.Stdout, "\033c")

	time.Sleep(time.Second)
	fmt.Fprint(os.Stdout, "hi")

	time.Sleep(time.Second)
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	if err := clear.Run(); err != nil {
		log.Fatal(err)
	}

	Hi()
	HiDup()
	Hello()

	time.Sleep(time.Second)
	Clear()
	fmt.Println("after clear")
}

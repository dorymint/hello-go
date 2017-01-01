package main

import (
	"fmt"
	"log"
	"os/exec"
)

// exec test

func main() {
	execName()
}

func execName() {
	fmt.Println("test2")
	fmt.Println(exec.LookPath("ls"))
	exe := exec.Command("ls")

	exist, err := exe.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(exist))

	exe = exec.Command("ls", "../../")
	exist2, err := exe.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(exist2))
}

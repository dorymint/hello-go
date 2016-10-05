package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)


func main() {
	log.Println("main")
	cmd := exec.Command("echo", "hello")
	fmt.Println("cmd show")
	fmt.Printf("%T\n", cmd)
	fmt.Println(cmd.Args)
	fmt.Println(cmd.Dir)
	fmt.Println(cmd.Env)
	fmt.Println(cmd.ExtraFiles)
	fmt.Println(cmd.Path)
	fmt.Println(cmd.ProcessState)
	fmt.Println(cmd.Stderr)
	fmt.Println(cmd.Stdin)
	fmt.Println(cmd.Stdout)
	fmt.Println(cmd.SysProcAttr)

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\ntodogotcha")

	todogotcha := exec.Command("todogotcha", "-root", "./", "-result", "off")
	// TODO: This TODO is test.
	todogotcha.Stdout = os.Stdout
	todogotcha.Stderr = os.Stderr
	fmt.Println("todogotcha args=",todogotcha.Args)
	fmt.Println("todogotcha path=",todogotcha.Path)

	if err := todogotcha.Run(); err != nil {
		log.Fatal(err)
	}
}


package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("u:", u)
	fmt.Println("GroupIds:", func() string { return fmt.Sprintln(u.GroupIds()) }())
	fmt.Println("Gid:", u.Gid)
	fmt.Println("HomeDir:", u.HomeDir)
	fmt.Println("Name:", u.Name)
	fmt.Println("Uid:", u.Uid)
	fmt.Println("Username:", u.Username)
}

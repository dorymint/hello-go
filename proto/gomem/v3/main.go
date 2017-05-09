package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gomem"
	"subcmd"
)

// read
func read(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		panic(sc.Err())
	}
	return sc.Text()
}

func gomemNew() (string, error) {
	g, err := gomem.New(strings.TrimSpace(read("new filepath:>")), true)
	if err != nil {
		return "", err
	}
	g.Title = read("title:>")
	g.Content = read("content:>")
	return g.WriteFile()
}

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("LOG:> ")

	gs := gomem.GomemsNew("./t")
	fmt.Println(gs)

	sub := subcmd.NewWithExit()
	sub["new"] = gomemNew
	fmt.Println("run sub.Repl")
	if err := sub.Repl("repl test:>"); err != nil {
		log.Fatal(err)
	}
}

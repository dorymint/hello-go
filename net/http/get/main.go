package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// Example: godoc -http ":6060"
	url = flag.String("url", "http://localhost:6060/", "")
	out      = flag.String("out", "", "")
)

// Checking after persing flags
func argsCheck() {
	if len(flag.Args()) != 0 {
		fmt.Printf("cmd = %v\n", os.Args)
		fmt.Println("-----| Unknown option |-----")
		for _, x := range flag.Args() {
			fmt.Println(x)
		}
		fmt.Println("-----| Usage |-----")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// simple confirm
// Exmaple: if confirm(os.Stdin, "exit[yes:no]>", 2) { os.Exit(0) }
func confirm(stdin *os.File, msg string, count int) bool {
	fmt.Print(msg)
	for sc, i := bufio.NewScanner(stdin), 0; sc.Scan() && i < count; i++ {
		if sc.Err() != nil {
			log.Fatal(sc.Err())
		}
		switch sc.Text() {
		case "yes", "y":
			return true
		case "no", "n":
			return false
		default:
			fmt.Println(sc.Text())
			fmt.Print(msg)
		}
	}
	return false
}

func main() {
	flag.Parse()
	argsCheck()

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
		///return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		///return err
	}

	if *out != "" {
		if _, err := os.Stat(*out); err == nil {
			if confirm(os.Stdin, *out+": override? [yes:no]>", 2) == false {
				log.Fatal("do not override")
				///return fmt.Errorf("do not override")
			}
		}
		if err := ioutil.WriteFile(*out, b, 0600); err != nil {
			log.Fatal(err)
			///return err
		}
	} else {
		fmt.Print(string(b))
	}
}

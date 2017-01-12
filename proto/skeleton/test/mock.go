// mock for era
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := proc(); err != nil {
		log.Fatal(err)
	}
}

func proc() error {
	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			return fmt.Errorf("proc: %s", sc.Err())
		}
		switch txt := sc.Text(); txt {
		default:
			if i, err := strconv.Atoi(txt); err != nil {
				log.Println(err)
				continue
			} else {
				log.Println(txt, i, []byte(strconv.Itoa(i)), string(strconv.Itoa(i)))
			}
		case "exit":
			return nil
		case "numtest":
			fmt.Println(txt)
			fmt.Println(1, []byte(strconv.Itoa(1)), string(strconv.Itoa(1)))
			fmt.Println(2, []byte(strconv.Itoa(2)), string(strconv.Itoa(2)))
			fmt.Println(3, []byte(strconv.Itoa(3)), string(strconv.Itoa(3)))
		}
	}
	return nil
}

// pkg/encoding/json/interaceive.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func run() error {
	var v interface{}
	var err error
	sc := bufio.NewScanner(os.Stdin)
	scan := func() bool {
		fmt.Print("input jSON(exit: <C-c> Or <C-d>) >")
		return sc.Scan()
	}
	for scan() {
		err = json.Unmarshal([]byte(sc.Text()), &v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%+v\n", v)
		v = nil
	}
	if err = sc.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

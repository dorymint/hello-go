package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

const contents = `hello,0,"string, string"
world,100,"string"`

func main() {
	r := csv.NewReader(strings.NewReader(contents))
	fmt.Println(r.Read())
	fmt.Println(r.ReadAll())
}

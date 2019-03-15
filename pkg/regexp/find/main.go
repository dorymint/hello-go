// find.
package main

import (
	"fmt"
	"regexp"
)

func main() {
	validate := regexp.MustCompile("^/(edit)/?(file)?$")
	tests := []string{
		"/edit",
		"/edit/",
		"/edit/file",
		"http://edit",
		"https://edit",
	}

	for _, x := range tests {
		fmt.Printf("\tin:%s\n", x)
		fmt.Printf("\tout:%s\n", validate.FindString(x))
		fmt.Printf("\tin:%s\n", x)
		fmt.Printf("\tout:%q\n\n", validate.FindStringSubmatch(x))
	}
}

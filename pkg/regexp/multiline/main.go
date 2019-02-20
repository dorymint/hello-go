package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	str := `abc def
abc
def
(abc)

m
?
?m
(?m)
`
	// (?m) is flag of multi-line
	match := regexp.MustCompile("(?m)^abc$")

	fmt.Printf("src:%V\n\n", str)

	fmt.Printf("FindSTring:%V\n\n", match.FindString(str))

	index := match.FindStringIndex(str)
	fmt.Printf("FindStringIndex:%V\n\n", index)
	fmt.Printf("str[:index[0]]:%V\n\n", str[:index[0]])
	fmt.Printf("str[index[1]:]:%V\n\n", str[index[1]:])

	r := bytes.NewBufferString(str)
	fmt.Printf("FindIndex:%V\n\n", match.FindIndex([]byte(str)))
	ri := match.FindReaderIndex(r)
	fmt.Printf("FindReaderIndex:%V\n\n", ri)
	fmt.Printf("str == r.String():%V\n\n", str == r.String())
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read at=%v\n", string(b[ri[1]:]))
	}
	fmt.Printf("FindStringSubMatch:%V\n\n", match.FindStringSubmatch(str))

	fmt.Printf("Split:%V\n%V\n\n", match.Split(str, 2))
}

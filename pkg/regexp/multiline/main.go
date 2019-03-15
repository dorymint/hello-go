// multiline.
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

	fmt.Printf("src:%+v\n\n", str)

	fmt.Printf("FindSTring:%+v\n\n", match.FindString(str))

	index := match.FindStringIndex(str)
	fmt.Printf("FindStringIndex:%+v\n\n", index)
	fmt.Printf("str[:index[0]]:%+v\n\n", str[:index[0]])
	fmt.Printf("str[index[1]:]:%+v\n\n", str[index[1]:])

	r := bytes.NewBufferString(str)
	fmt.Printf("FindIndex:%+v\n\n", match.FindIndex([]byte(str)))
	ri := match.FindReaderIndex(r)
	fmt.Printf("FindReaderIndex:%+v\n\n", ri)
	fmt.Printf("str == r.String():%+v\n\n", str == r.String())
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read at=%v\n", string(b[ri[1]:]))
	}
	fmt.Printf("FindStringSubMatch:%+v\n\n", match.FindStringSubmatch(str))

	fmt.Printf("Split:%+v\n%+v\n\n", match.Split(str, 2))
}

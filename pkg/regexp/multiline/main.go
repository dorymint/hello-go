package main

import (
	"fmt"
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
	match := regexp.MustCompile("(?m)^(abc)$")

	fmt.Printf("src:%V\n\n", str)
	fmt.Printf("fs:%V\n\n", match.FindString(str))
	index := match.FindStringIndex(str)
	fmt.Printf("fsi:%V\n\n", index)
	fmt.Printf("fss:%V\n\n", match.FindStringSubmatch(str))

	fmt.Printf("str[:index[0]]:%V\n", str[:index[0]])
	fmt.Printf("str[index[1]:]:%V\n", str[index[1]:])
}

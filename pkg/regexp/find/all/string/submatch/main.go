// FindAllStringSubmatch.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	file := flag.String("file", "", "source file")
	src := flag.String("src", "", "provide source string")
	flag.Parse()
	switch flag.NArg() {
	case 0:
		// pass
	case 1:
		if *file == "" {
			*file = flag.Arg(0)
			break
		}
		fallthrough
	default:
		flag.Usage()
		panic("unexpected arguments")
	}

	if *src != "" && *file != "" {
		panic("duplicate specification")
	}

	if *file != "" {
		b, err := ioutil.ReadFile(*file)
		if err != nil {
			panic(err)
		}
		*src = string(b)
	}

	if *src == "" {
		panic("not specified source")
	}

	var (
		sc   = bufio.NewScanner(os.Stdin)
		scan = func() bool {
			fmt.Print("pat$")
			return sc.Scan()
		}
	)
	for scan() {
		pat := sc.Text()
		re, err := regexp.Compile(pat)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[err]:%v:%q\n", err, pat)
			continue
		}
		result := re.FindAllStringSubmatch(*src, -1)
		fmt.Printf("Result:\n%+v\n", result)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
}

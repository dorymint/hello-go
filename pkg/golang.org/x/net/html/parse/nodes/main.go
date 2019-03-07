package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/net/html"
)

var ntm = map[html.NodeType]string{
	html.ErrorNode:    "ErrorNode",
	html.TextNode:     "TextNode",
	html.DocumentNode: "DocumentNode",
	html.ElementNode:  "ElementNode",
	html.CommentNode:  "CommentNode",
	html.DoctypeNode:  "DoctypeNode",
}

var file string

func init() {
	flag.StringVar(&file, "html", filepath.Join("testdata", "test.html"), "path to html")
}

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	n, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	fmt.Printf("*html.Node:%T\n\n", n)

	var s string
	var f func(*html.Node)
	f = func(n *html.Node) {
		s += fmt.Sprintf("// *html.Node:%+v\n", n)
		s += fmt.Sprintf("n.Type:%v\n", ntm[n.Type])
		s += fmt.Sprintf("n.Data:%#v\n", n.Data)
		s += fmt.Sprintf("n.DataAtom:%v\n", n.DataAtom)
		s += fmt.Sprintf("n.Namespace:%v\n", n.Namespace)
		s += fmt.Sprintf("n.Attr:\n")
		if len(n.Attr) != 0 {
			for _, attr := range n.Attr {
				s += fmt.Sprintf("\t%+v\n", attr)
			}
		}
		s += "\n"

		// tail recursion
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	fmt.Println(s)
}

package main

import (
	"bytes"
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

func main() {
	testdata := filepath.Join("testdata", "test1.html")
	b, err := ioutil.ReadFile(testdata)
	if err != nil {
		panic(err)
	}

	n, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	fmt.Printf("*html.Node:%T\n", n)
	fmt.Printf("*html.Node:%+v\n\n", n)

	var f func(*html.Node)
	f = func(n *html.Node) {
		fmt.Printf("// *html.Node:%v\n", n)
		s := fmt.Sprintf("n.Type:%v\n", ntm[n.Type])
		if len(n.Attr) != 0 {
			s += fmt.Sprintf("\tn.Attr:\n")
			for _, attr := range n.Attr {
				s += fmt.Sprintf("\t\t%v\n", attr)
			}
		}
		s += fmt.Sprintf("\tn.Data:%v\n", n.Data)
		s += fmt.Sprintf("\tn.DataAtom:%v\n", n.DataAtom)
		s += "\n"
		fmt.Println(s)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
}

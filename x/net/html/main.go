package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// REF: godoc.org/golang.org/x/net/html

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func getByte(file string) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func tokenize(b []byte) error {
	z := html.NewTokenizer(bytes.NewReader(b))
	for i := 0; ; i++ {
		split("count: " + strconv.Itoa(i))
		tt := z.Next()
		if tt == html.ErrorToken {
			return z.Err()
		}
		fmt.Println("tt.String       :", tt.String())
		fmt.Println("z.Raw           :", string(z.Raw()))
		fmt.Println("z.Text          :", string(z.Text()))
		fmt.Println("z.Token.String  :", z.Token().String())
		fmt.Println("z.Token.Type    :", z.Token().Type)
	}
}

func tokenTypes() {
	fmt.Println(html.ErrorToken)
	fmt.Println(html.TextToken)
	fmt.Println(html.SelfClosingTagToken)
	fmt.Println(html.StartTagToken)
	fmt.Println(html.EndTagToken)
	fmt.Println(html.CommentToken)
	fmt.Println(html.DoctypeToken)

	fmt.Println(html.ErrorNode)
	fmt.Println(html.TextNode)
	fmt.Println(html.DocumentNode)
	fmt.Println(html.ElementNode)
	fmt.Println(html.CommentNode)
	fmt.Println(html.DoctypeNode)
}

func tokenizeSwitch(b []byte) error {
	depth := 0
	z := html.NewTokenizer(bytes.NewReader(b))
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() != io.EOF {
				return z.Err()
			}
			return nil
		case html.TextToken:
			if depth > 0 {
				fmt.Println(string(z.Text()))
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
}

func parse(b []byte) error {
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		return err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("", a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}

func q1(f string) {
	split("q1 " + f)
	b := getByte(f)
	split("tokenize")
	if err := tokenize(b); err != nil && err != io.EOF {
		log.Fatal("tokenize: " + err.Error())
	}
	split("tokenizeSwitch")
	if err := tokenizeSwitch(b); err != nil {
		log.Fatal(err)
	}
	split("parse")
	if err := parse(b); err != nil {
		log.Fatal(err)
	}
}

func q2(f string) {
	b := getByte(f)
	z := html.NewTokenizer(bytes.NewReader(b))

	for tt := z.Next(); ; tt = z.Next() {
		if tt == html.ErrorToken {
			fmt.Fprintln(os.Stderr, z.Err())
			return
		}
		fmt.Println("tt   :", tt)
		fmt.Println("z.Raw:", string(z.Raw()))
	}
}

func parseContent(file string) {
	var check []string

	b := getByte(file)
	node, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "meta" {
			ogImage := false
			for _, a := range n.Attr {
				fmt.Println(" ", a.Key)
				fmt.Println("   ", a.Val)
				if ogImage && a.Key == "content" {
					check = append(check, a.Val)
				}
				if a.Key == "property" && a.Val == "og:image" {
					ogImage = true
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
	for _, c := range check {
		fmt.Println("check:", c)
	}
}

func main() {
	split("tokenTypes")
	tokenTypes()
	if false {
		split("q1")
		q1("./t.html")
		q1("./mock.html")
		split("q2")
		q2("./mock.html")
	}
	split("parseContent")
	parseContent("./mock.html")
}

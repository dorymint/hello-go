/*

parse html

	getattr -html /path/file.html -json '{
		"elem": "a",
		"attr": "href",
		"want_attr": {
			"class": "id"
		}
	}'

	getattr -file /path/html -config /path/config.json

*/
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

type Filter struct {
	Elem string `json:"elem"`
	Attr string `json:"attr"`

	// map[key]val
	WantAttrMap map[string]string `json:"want_attr"`
}

func NewFilter() *Filter {
	return &Filter{
		Elem:        "",
		Attr:        "",
		WantAttrMap: make(map[string]string),
	}
}

func (fil *Filter) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &fil)
}

func (fil *Filter) ReadJSON(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return fil.Unmarshal(b)
}

func (fil *Filter) WriteJSON(w io.Writer) error {
	b, err := json.MarshalIndent(fil, "", "  ")
	if err != nil {
		return err
	}
	_, err = io.Copy(w, bytes.NewReader(b))
	return err
}

func (fil *Filter) getAttr(attrs []html.Attribute) *html.Attribute {
	var (
		attr  *html.Attribute
		valid = true
	)
	var (
		want = len(fil.WantAttrMap)
		done = 0
	)
	for i, a := range attrs {
		if fil.Attr == a.Key {
			attr = &attrs[i]
			continue
		}

		wantval, ok := fil.WantAttrMap[a.Key]
		if !ok {
			continue
		}
		if valid {
			valid = wantval == a.Val
		}
		done++
	}
	if attr != nil && valid && want == done {
		return attr
	}
	return nil
}

func (fil *Filter) getAttrVals(node *html.Node) (vals []string) {
	var crawl func(*html.Node)
	crawl = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode && n.Data == fil.Elem {
			if a := fil.getAttr(n.Attr); a != nil {
				vals = append(vals, a.Val)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawl(c)
		}
	}
	crawl(node)
	return vals
}

func (fil *Filter) GetAttrVals(fileHTML string) ([]string, error) {
	b, err := ioutil.ReadFile(fileHTML)
	if err != nil {
		return nil, err
	}
	n, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return fil.getAttrVals(n), nil
}

func WriteTemplate(w io.Writer) error {
	fil := NewFilter()
	fil.Elem = "a"
	fil.Attr = "href"
	fil.WantAttrMap["class"] = "preview"
	return fil.WriteJSON(w)
}

var defaultHTML = filepath.Join("testdata", "test.html")
var defaultConfig = filepath.Join("testdata", "config.json")

var opt struct {
	file       string
	configJSON string
	stringJSON string
	template   bool
}

func init() {
	flag.StringVar(&opt.file, "html", defaultHTML, "path to html")
	flag.StringVar(&opt.configJSON, "config", defaultConfig, "path to JSON format config")
	flag.StringVar(&opt.stringJSON, "json", "", "direct specify")
	flag.BoolVar(&opt.template, "template", false, "output JSON config template")
}

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		panic("invalid arguments")
	}

	if opt.template {
		if err := WriteTemplate(os.Stdout); err != nil {
			panic(err)
		}
		return
	}

	fil := NewFilter()

	if opt.configJSON != "" {
		if err := fil.ReadJSON(opt.configJSON); err != nil {
			panic(err)
		}
	}
	// marge
	if opt.stringJSON != "" {
		if err := fil.Unmarshal([]byte(opt.stringJSON)); err != nil {
			panic(err)
		}
	}

	vals, err := fil.GetAttrVals(opt.file)
	if err != nil {
		panic(err)
	}
	if len(vals) != 0 {
		fmt.Printf("%s\n", strings.Join(vals, "\n"))
	}
}

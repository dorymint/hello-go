package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var DIR = filepath.Join("testdata")

func TestGetAttrVals(t *testing.T) {
	type file struct {
		html   string
		config string
		exp    string
	}
	var files []file
	fis, err := ioutil.ReadDir(DIR)
	if err != nil {
		t.Fatal(err)
	}
	ndir := 0
	for _, fi := range fis {
		if fi.IsDir() {
			ndir++
		}
	}
	for i := 0; i != ndir; i++ {
		dir := filepath.Join(DIR, fmt.Sprintf("%d", i))
		var f file
		f.html = filepath.Join(dir, "test.html")
		f.config = filepath.Join(dir, "config.json")
		f.exp = filepath.Join(dir, "exp.txt")
		files = append(files, f)
	}

	for _, test := range files {
		fil := NewFilter()
		if err := fil.ReadJSON(test.config); err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadFile(test.exp)
		if err != nil {
			t.Fatal(err)
		}
		exp := strings.Fields(string(b))

		vals, err := fil.GetAttrVals(test.html)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(exp, vals) {
			t.Errorf("exp:%+v but out:%+v", exp, vals)
		}
	}
}

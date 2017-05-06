package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TempGomem(content string) (string, error) {
	f, err := ioutil.TempFile("", "gomem")
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return "", err
	}
	name := f.Name()
	return name, nil
}

func TestGomem(t *testing.T) {
	data := []struct {
		ext *Gomem
		out *Gomem
	}{
		{
			ext: &Gomem{Title: "title", Content: "content", Tags: []string{}},
			out: GomemNew("title", "content", nil),
		},
		{
			ext: &Gomem{Title: "title", Content: "content", Tags: []string{}},
			out: GomemNew("title", "content", []string{}),
		},
		{
			ext: &Gomem{Title: "title", Content: "content", Tags: []string{"tags", "tag"}},
			out: GomemNew("title", "content", []string{"tags", "tag"}),
		},
	}

	for i, v := range data {
		if !reflect.DeepEqual(v.ext, v.out) {
			t.Fatalf("count %d:\next: %q\nout: %q\n", i, v.ext, v.out)
		}
	}
}

func TestReadFile(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(pwd)
}

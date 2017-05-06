package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Gomem struct {
	Title   string
	Content string
	Tags    []string
}

type Gomems map[string]*Gomem

func GomemNew(title string, content string, tags []string) *Gomem {
	if tags == nil {
		tags = []string{}
	}
	return &Gomem{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
}

func (g *Gomem) String() string {
	if g == nil {
		return ""
	}
	b, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (gs *Gomems) String() string {
	if (*gs) == nil {
		return ""
	}
	b, err := json.MarshalIndent(gs, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (gs *Gomems) Add(fpath string, g *Gomem) error {
	if (*gs) == nil {
		panic("main.go: *gs Add(): *gs == nil")
	}
	if _, ok := (*gs)[fpath]; ok {
		return fmt.Errorf("gs.Add: fpath exists")
	}
	(*gs)[fpath] = g
	return nil
}

func (gs *Gomems) AddFile(fpath string, title string, tags []string,  override bool) error {
	if (*gs) == nil {
		return fmt.Errorf("gs.AddFile: gs == nil ")
	}
	if _, ok := (*gs)[fpath]; ok && override == false {
		return fmt.Errorf("gs.AddFile: fpath exists")
	}

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	(*gs)[fpath] = GomemNew(title, string(b), tags)
	return nil
}

///
/// TODO: implementation Gomems shell
///
func mock(g *Gomems) error {
	fmt.Println("called mock")
	return nil
}

type command struct {
	name string
	arg  string
	f    func(*Gomems) error
}

var commands map[string]command

func repl(g *Gomems) error {
	c, ok := commands["get"]
	if !ok {
		return fmt.Errorf("invalid command")
	}
	err := c.f(g)
	for err == nil {
	}
	return err
}

func main() {
	g := make(Gomems)
	g["hello"] = GomemNew("hi", "test", []string{"tags"})
	g["world"] = GomemNew("workd", "nyan", []string{"new"})
	g["nyan"] = GomemNew("", "", nil)
	fmt.Println(&g)
	fmt.Println(g["hello"])
	data := fmt.Sprintln(&g)
	fmt.Println("data1: ", data)
	data = fmt.Sprintln(g)
	fmt.Println("data2: ", data)

	if err := repl(&g); err != nil {
		log.Fatal(err)
	}
}

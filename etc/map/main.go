package main

import (
	"fmt"
	"log"
)

type Gomem struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Gomems map[string]*Gomem

func gomemNew(title, content string) *Gomem {
	return &Gomem{
		Title:   title,
		Content: content,
	}
}

func (g *Gomems) mock() error {
	if g == nil {
		return fmt.Errorf("g == nil")
	}
	return fmt.Errorf("g != nil: test error")
}

func main() {
	log.SetFlags(log.Lshortfile)
	const testPath = "./test.json"

	gs := make(Gomems)
	fmt.Println(gs)
	gs[testPath] = gomemNew("test", "content")
	fmt.Println(gs)
	fmt.Println(gs[testPath])
	fmt.Println(gs.mock())
}

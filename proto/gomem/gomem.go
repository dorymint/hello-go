package gomem

import (
	"encoding/json"
	"log"
)

type Gomem struct {
	Title   string
	Content string
	Tags    []string
}

type Gomems map[string]*Gomem

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

func (g *Gomems) String() string {
	if (*g) == nil {
		return ""
	}
	b, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

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

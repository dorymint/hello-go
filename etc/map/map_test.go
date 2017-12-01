package m

import (
	"fmt"
	"testing"
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

func TestGomem(t *testing.T) {
	const testPath = "./test.json"

	gs := make(Gomems)
	t.Log(gs)
	gs[testPath] = gomemNew("test", "content")
	t.Log(gs)
	t.Log(gs[testPath])
	t.Log(gs.mock())
}

func TestMap(t *testing.T) {
	m := make(map[rune]bool)
	t.Log("m len:", len(m))
	m['h'] = true
	t.Log("m+\"h\" len:", len(m))

	m2 := make(map[rune]bool, 2)
	t.Log("m2 len:", len(m2))
	m2['h'] = true
	t.Log("m2+\"h\" lne:", len(m2))

	for _, s := range("hello world") {
		m2[s] = true
	}
	t.Log("m2+\"hello world\"", len(m2))
}

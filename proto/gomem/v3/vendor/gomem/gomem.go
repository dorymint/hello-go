package gomem

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Gomem JSON structure
type Gomem struct {
	Title   string `json:"title"`
	Content string `json:"content"`

	filepath string
	override bool
}

// New return *Gomem init filepath and override flag
// if fpath is exists and write is false then return error
func New(fpath string, write bool) (*Gomem, error) {
	if _, err := os.Stat(fpath); err == nil && write != true {
		return nil, errors.New("New: file exists: " + fpath)
	}
	return &Gomem{filepath: fpath, override: write}, nil
}

// ReadFile load from g.filepath
func (g *Gomem) ReadFile() (string, error) {
	if g == nil {
		return "", errors.New("*Gomem.ReadFile: g == nil")
	}
	b, err := ioutil.ReadFile(g.filepath)
	if err != nil {
		return "", errors.New("*Gomem.ReadFile " + g.filepath + ": " + err.Error())
	}
	err = json.Unmarshal(b, g)
	if err != nil {
		return "", errors.New("*Gomem.ReadFile " + g.filepath + ": " + err.Error())
	}
	return "file readed: " + g.filepath, nil
}

// WriteFile write to g.filepath
func (g *Gomem) WriteFile() (string, error) {
	if g == nil {
		return "", errors.New("*Gomem.WriteFile: g == nil")
	}
	if _, err := os.Stat(g.filepath); err == nil && g.override != true {
		return "", errors.New("*Gomem.WriteFile " + g.filepath + ": file exitsts")
	}
	b, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "", errors.New("*Gomem.WriteFile " + g.filepath + err.Error())
	}
	if err := ioutil.WriteFile(g.filepath, b, 0600); err != nil {
		return "", errors.New("g.WriteFile: " + err.Error())
	}
	return g.filepath + ": writed", nil
}


// Gomems map of Gomem and data directory
type Gomems struct {
	Gmap map[string]*Gomem
	dir  string // reconsider base directory
}

// GomemsNew read from dir return map for Gomem
func GomemsNew(dir string) Gomems {
	gs := Gomems{
		Gmap: make(map[string]*Gomem),
		dir:  dir,
	}
	if _, err := gs.IncludeJSON(); err != nil {
		panic(err)
	}
	return gs
}

// IncludeJSON include from Gomems.dir
// mapping dir+basename => *Gomem
func (gs *Gomems) IncludeJSON() (string, error) {
	if gs == nil || gs.Gmap == nil {
		return "", errors.New("*Gomems.getJSON: nil error")
	}
	bases, err := getBaseJSON(gs.dir)
	if err != nil {
		return "", err
	}
	for _, x := range bases {
		key := filepath.Join(gs.dir, x)
		if g, ok := gs.Gmap[key]; ok {
			if _, err := g.ReadFile(); err != nil {
				return "", err
			}
			continue
		}
		g := &Gomem{filepath: key}
		gs.Gmap[key] = g
		if _, err := g.ReadFile(); err != nil {
			return "", err
		}
	}
	return "getJSON: JSON files included from " + gs.dir, nil
}


// return json files base name in dir
func getBaseJSON(dir string) ([]string, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var str []string
	for _, info := range infos {
		if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".json") {
			str = append(str, info.Name())
		}
	}
	return str, nil
}

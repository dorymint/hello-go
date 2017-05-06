package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Gomem JSON structure
type Gomem struct {
	Title   string `json:"title"`
	Content string `json:"content"`

	filepath string
	override bool
}

// Gomems map of Gomem and data directory
type Gomems struct {
	Gmap map[string]*Gomem
	Dir  string
}

// ReadFile load from g.filepath
func (g *Gomem) ReadFile() (string, error) {
	if g == nil {
		return "", fmt.Errorf("*Gomem.ReadFile: g == nil")
	}
	b, err := ioutil.ReadFile(g.filepath)
	if err != nil {
		return "", fmt.Errorf("*Gomem.ReadFile: %q", err)
	}
	err = json.Unmarshal(b, g)
	if err != nil {
		return "", fmt.Errorf("*Gomem.ReadFile: %q", err)
	}
	return g.filepath + ": readed", nil
}

// WriteFile write to g.filepath
func (g *Gomem) WriteFile() (string, error) {
	if g == nil {
		return "", fmt.Errorf("*Gomem.OutFile: g == nil")
	}
	if _, err := os.Stat(g.filepath); err == nil && g.override != true {
		return "", fmt.Errorf("*Gomem.OutFile: %s is exists", g.filepath)
	}

	b, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(g.filepath, b, 0600); err != nil {
		return "", fmt.Errorf("g.OutFile: %q", err)
	}
	return g.filepath + ": writed", nil
}

// test func
func (g *Gomem) echo() (string, error) {
	fmt.Println("echo.tilte=", g.Title)
	fmt.Println("echo.content=", g.Content)
	g.Title = "after"
	return "title changed", nil
}

// impl repl
// make subcommands

// SubCommands interp functions for Repl
type SubCommands map[string]func() (string, error)

// ErrValidExit for valid exit, use Repl, return nil
var ErrValidExit = errors.New("valid exit")

func exit() (string, error) {
	return "", ErrValidExit
}

// Repl Read Eval P.*? Loop
// call function in SubCommands[string]
// string is from os.Stdin
// if return ErrValidExit then return nil
func (sub SubCommands) Repl(prefix string) error {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prefix)
		if !sc.Scan() {
			return fmt.Errorf("fail sc.Scan")
		}
		if sc.Err() != nil {
			return sc.Err()
		}
		f, ok := sub[strings.TrimSpace(sc.Text())]
		if !ok {
			fmt.Printf("invalid subcommand: %q\n", sc.Text())
			continue
		}
		result, err := f()
		if err != nil {
			switch err {
			case ErrValidExit:
				return nil
			default:
				return err
			}
		}
		fmt.Println(result)
	}
}

// SubNew return SubCommands with included interpreter for exit
// "exit" "quit" ":q"
// them are call the valid exit, return nil
func SubNew() SubCommands {
	sub := make(SubCommands)
	sub["exit"] = exit
	sub["quit"] = exit
	sub[":q"] = exit
	return sub
}

func GomemsNew(dir string) Gomems {
	gs := Gomems{
		Gmap: make(map[string]*Gomem),
		Dir:  dir,
	}
	_, err := gs.getJSON()
	if err != nil {
		panic(err)
	}
	return gs
}

// read
func read(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		panic(sc.Err())
	}
	return sc.Text()
}

// RECONSIDER:
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

// RECONSIDER:
func (gs *Gomems) getJSON() (string, error) {
	if gs == nil || gs.Gmap == nil {
		return "", fmt.Errorf("*Gomems.getJSON: nil error")
	}
	bases, err := getBaseJSON(gs.Dir)
	if err != nil {
		return "", err
	}
	for _, x := range bases {
		_, ok := gs.Gmap[gs.Dir+x]
		if ok {
			continue
		}
		gs.Gmap[gs.Dir+x] = &Gomem{
			filepath: gs.Dir + x,
		}
		_, err := gs.Gmap[gs.Dir+x].ReadFile()
		if err != nil {
			return "", err
		}
	}
	return "getJSON: JSON files included from " + gs.Dir, nil
}

// flags
var (
	readdir = flag.String("dir", "./", "")
	defjson = flag.String("json", "./gomem.json", "")
	cdhome = flag.String("home", "./t", "")
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}
func main() {
	flag.Parse()
	if err := os.Chdir(*cdhome); err != nil {
		panic(err)
	}

	log.SetFlags(log.Lshortfile)
	log.SetPrefix("LOG:> ")

	var g Gomem
	var gs Gomems
	var result string
	var err error

	split("gomem.json")
	g.filepath = "./gomem.json"
	fmt.Printf("g=%q\n", g)
	result, err = g.ReadFile()
	log.Println(result, err)
	fmt.Printf("AFTER ReadFile:\ng=%q\n", g)

	split("gomem_alter.json")
	g.filepath = "./gomem_alter.json"
	result, err = g.WriteFile()
	log.Println(result, err)
	b, err := ioutil.ReadFile("./gomem_alter.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AFTER ioutil.ReadFile:\n", string(b))
	str, err := getBaseJSON("./")
	fmt.Println("getBaseJSON str&err=", str, err)

	split("gs")
	fmt.Println("gs=", gs)
	gs.Dir = "./"
	gs.Gmap = make(map[string]*Gomem)
	result, err = gs.getJSON()
	log.Println(result, err)
	fmt.Println("AFTER getJSON:\ngs=", gs)

	split("gs.Gmap")
	for key, x := range gs.Gmap {
		fmt.Println(key, x)
	}

	split("tg")
	tg := gs.Gmap["./gomem_alter.json"]
	fmt.Printf("tg=%q\n", tg)
	fmt.Println("tg.filepath=", tg.filepath)

	split("result")
	fmt.Println("gs:", gs, "\n", "g:", g, "\n", "tg:", tg)

	split("GomemsNew")
	gsnext := GomemsNew("./")
	fmt.Println("gsnext=", gsnext)
	for key, x := range gs.Gmap {
		fmt.Println(key, x)
	}
}

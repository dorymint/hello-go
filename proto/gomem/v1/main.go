package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/* Gomem */

// Gomem memo style json
type Gomem struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Gomems map[file path]*Gomem
type Gomems map[string]*Gomem

// AddFile read from fpath to add to g[fpath], if override true then do it
func (g *Gomems) AddFile(fpath string, override bool) error {
	if (*g) == nil {
		return fmt.Errorf("g.AddFile: *g == nil")
	}
	if _, ok := (*g)[fpath]; ok && !override {
		return fmt.Errorf("g.AddFile: g[%s] is not nil", fpath)
	}

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("g.AddFile: %q", err)
	}
	buf := new(Gomem)
	err = json.Unmarshal(b, buf)
	if err != nil {
		return fmt.Errorf("g.AddFile: %q", err)
	}
	(*g)[fpath] = buf
	return nil
}

// OutFile OutPut g[fpath] to fpath, if override true then do it
func (g *Gomems) OutFile(fpath string, override bool) error {
	if (*g) == nil {
		return fmt.Errorf("g.OutFile: *g == nil")
	}
	v, ok := (*g)[fpath]
	if !ok {
		return fmt.Errorf("g.OutFile: g[%s] is not found", fpath)
	}
	if _, err := os.Stat(fpath); err == nil && !override {
		return fmt.Errorf("g.OutFile: %s is exists", fpath)
	}

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fpath, b, 0600)
}

func (g *Gomems) String() string {
	if (*g) == nil {
		return ""
	}
	var str string
	for key, v := range *g {
		str += fmt.Sprintf("%q %q\n", key, v)
	}
	return str
}

/* Flags */

// Flags structuer of flag
type Flags struct {
	defpath *string
	memdir  *string
}

var fs = Flags{
	defpath: flag.String("gomem", "./gomem.json", ""),
	memdir:  flag.String("memdir", "./", ""),
}

func (fs *Flags) init() {
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		log.Fatal("invalid: ", flag.Args())
	}
}

func getJSON(dir string) ([]string, error) {
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

// cmd: gomem
//    : 0. parse flags
//    : 1. load the default data directory
//    : 2. search suffix ".json" to add Gomems
//    : 3. switch interactive noninteractive
// inter: 0. init to doit run repl
// noninter: 0. do the subcommands
// NEXT impl subcommands

/* read repl */
func read(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		panic(sc.Err())
	}
	return sc.Text()
}
func repl(msg string) error {
	fmt.Print(msg)
	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			return sc.Err()
		}
		switch s := strings.TrimSpace(sc.Text()); s {
		case "exit":
			return nil
		case "echo":
			fmt.Println(read("echo:>"))
		default:
			fmt.Println(s)
		}
		fmt.Print(msg)
	}
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("ERROR [gomem]: ")
	paths := []string{
		"./test.json",
	}

	// test Gomems
	gs := make(Gomems)
	gs[paths[0]] = &Gomem{"Gomems Test", "CONTENT"}
	fmt.Println(gs)
	fmt.Println(gs[paths[0]])

	// test OutFile
	if err := gs.OutFile(paths[0], true); err != nil {
		log.Println(err)
	}

	// test AddFile
	if err := gs.AddFile("./add_test.json", false); err != nil {
		log.Println(err)
	}
	fmt.Println("AFTER ADD:", gs)
	// Fail test
	gs["./dir_test.json"] = &Gomem{"DIR TEST", "EXPECT FAIL"}
	err := gs.OutFile("./dir_test.json", true)
	if err != nil {
		log.Println(err)
	}

	// test &gs.String()
	fmt.Println(&gs)
	gs["./new.json"] = &Gomem{"NEW", "NEWC"}
	if err := gs.OutFile("./new.json", true); err != nil {
		log.Println(err)
	}

	// test fs
	fs.init()
	fmt.Println(*fs.defpath)
	fmt.Println(*fs.memdir)
	jsons, err := getJSON(*fs.memdir)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(jsons)
	repl("gomem:>")
}

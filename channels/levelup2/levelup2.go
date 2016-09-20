package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// flags
var (
	root = flag.String("root", "./", "Specify search root")
)

func init() {
	var err error
	flag.Parse()
	*root, err = filepath.Abs(*root)
	fatalIF("init", err)
}

// str is from function infomation.
func fatalIF(str string, err error) {
	if err != nil {
		log.Fatalf("%s:%v\n", str, err)
	}
}

// Use wait group!!
func useWaitGroupCrawl(root string) ([]string, map[string][]os.FileInfo) {
	var Result []string
	DirsCache := make(map[string]bool)
	Infomap := make(map[string][]os.FileInfo)
	Wg := new(sync.WaitGroup)
	Mux := new(sync.Mutex)

	var dirsCrawl func(string)
	dirsCrawl = func(dirname string) {
		defer Wg.Done()
		Mux.Lock()
		if DirsCache[dirname] {
			return
		}
		DirsCache[dirname] = true
		Mux.Unlock()

		f, err := os.Open(dirname)
		if err != nil {
			log.Printf("crawl:%v\n", err)
			return
		}
		defer f.Close()

		info, err := f.Readdir(0)
		if err != nil {
			log.Printf("crawl info:%v", err)
			return
		}
		Mux.Lock()
		Infomap[dirname] = info
		Mux.Unlock()
		for _, x := range info {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				Mux.Lock()
				Result = append(Result, tmp)
				Mux.Unlock()
				Wg.Add(1)
				go dirsCrawl(tmp)
			}
		}
	}
	Wg.Add(1)
	dirsCrawl(root)
	Wg.Wait()
	return Result, Infomap
}
func crawltset() {
	dirslist, infomap := useWaitGroupCrawl(*root)
	for _, x := range dirslist {
		fmt.Println(x)
	}
	fmt.Println("root = ", *root)
	fmt.Println("length", len(dirslist))

	var fileCount int
	for _, x := range dirslist {
		for _, y := range infomap[x] {
			if strings.HasSuffix(y.Name(), ".go")  || strings.HasSuffix(y.Name(), ".txt"){
				fmt.Println(filepath.Join(x, y.Name()))
				fileCount++
				time.Sleep(time.Millisecond)
			}
		}
	}
	fmt.Printf("all dirs = %v\nfile count = %v\n", len(dirslist),  fileCount)
}

func main() {
	crawltset()
}

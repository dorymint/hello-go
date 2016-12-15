package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type crawlLocker struct {
	Gophers   int
	dirsCache map[string]bool
	errlist  *[]string
	mux       sync.Mutex
}

// increment and decrement
func (g *crawlLocker) ToGopher(n int) {
	g.mux.Lock()
	g.Gophers = g.Gophers + n
	g.mux.Unlock()
}

// SetDirsCache
// if parameter s is already cache then return false
func (g *crawlLocker) SetDirsCache(s string) bool {
	g.mux.Lock()
	defer g.mux.Unlock()

	if g.dirsCache[s] {
		return false
	}
	g.dirsCache[s] = true
	return true
}

// set errlist
func (g *crawlLocker) pushErrorList(s string) {
	g.mux.Lock()
	*g.errlist = append(*g.errlist, s)
	g.mux.Unlock()
}

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

// dirsCrawl find like directory crawler!
/// NEXT:Benchmark
func dirsCrawl(dirname string, ch chan string, errlist *[]string) {
	// gophers is goroutine counter!
	locker := new(crawlLocker)
	locker.Gophers = int(1)
	locker.dirsCache = make(map[string]bool)
	locker.errlist = errlist

	var crawl func(string)
	crawl = func(dirname string) {
		defer func() {
			locker.mux.Lock()
			locker.Gophers--
			if locker.Gophers == 0 {
				close(ch)
			} else if locker.Gophers < 0 {
				panic("crawl:Cannot managed Gophers!")
			}
			locker.mux.Unlock()
		}()
		if !locker.SetDirsCache(dirname) {
			return
		}

		f, err := os.Open(dirname)
		if err != nil {
			locker.pushErrorList(fmt.Sprintf("crawl:%v\n", err))
			return
		}
		defer f.Close()

		info, err := f.Readdir(0)
		if err != nil {
			locker.pushErrorList(fmt.Sprintf("crawl info:%v\n", err))
			return
		}

		for _, x := range info {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				ch <- tmp
				locker.ToGopher(1)
				go crawl(tmp)
			}
		}
	}
	go crawl(dirname)
}
func processDirsCrawl() ([]string, []string) {
	ch := make(chan string, 10)
	errlist := new([]string)
	var dirs []string
	dirsCrawl(*root, ch, errlist)

	for x, ok := <-ch; ok; x, ok = <-ch {
		dirs = append(dirs, x)
	}
	return dirs, *errlist
}


// do not use structure test
/// NEXT:Benchmark
func hayabusaCrawl(root string, ch chan string, Errlist *[]string) {
	// gophers are goroutine counter!
	Gophers := int(1)
	DirsCache := make(map[string]bool)
	Mux := new(sync.Mutex)

	var crawl func(string)
	crawl = func(dirname string) {
		defer func() {
			Mux.Lock()
			defer Mux.Unlock()
			Gophers--
			if Gophers == 0 {
				close(ch)
			} else if Gophers < 0 {
				panic("crawl:Cannot managed Gophers!")
			}
		}()

		Mux.Lock()
		if DirsCache[dirname] {
			Mux.Unlock()
			return
		}
		DirsCache[dirname] = true
		Mux.Unlock()

		f, err := os.Open(dirname)
		if err != nil {
			Mux.Lock()
			*Errlist = append(*Errlist, fmt.Sprintf("crawl:%v\n", err))
			Mux.Unlock()
			return
		}
		defer f.Close()

		info, err := f.Readdir(0)
		if err != nil {
			Mux.Lock()
			*Errlist = append(*Errlist, fmt.Sprintf("crawl:%v\n", err))
			Mux.Unlock()
			return
		}
		for _, x := range info {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				ch <- tmp
				Mux.Lock()
				Gophers++
				Mux.Unlock()
				go crawl(tmp)
			}
		}
	}
	go crawl(root)
}
// test
func processHayabusa() ([]string, []string) {
	ch := make(chan string, 10)
	errlist := new([]string)
	var dirs []string
	hayabusaCrawl(*root, ch, errlist)

	for x, ok := <-ch; ok; x, ok = <-ch {
		dirs = append(dirs, x)
	}
	return dirs, *errlist
}



// do not used goroutine
func simpleDirsCrawl(dirname string, errlist *[]string) []string {
	dirsCache := make(map[string]bool)
	var result []string

	var crawl func(string)
	crawl = func(dirname string) {
		if dirsCache[dirname] {
			return
		}
		dirsCache[dirname] = true

		f, err := os.Open(dirname)
		if err != nil {
			*errlist = append(*errlist, fmt.Sprintf("crawl:%v\n", err))
			return
		}
		defer f.Close()

		info, err := f.Readdir(0)
		if err != nil {
			*errlist = append(*errlist, fmt.Sprintf("crawl:%v\n", err))
			return
		}
		for _, x := range info {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				result = append(result, tmp)
				crawl(tmp)
			}
		}
	}
	crawl(dirname)
	return result
}
func processSimpleCrawl() ([]string, []string) {
	errlist := new([]string)
	result := simpleDirsCrawl(*root, errlist)
	return result, *errlist
}


// Get Crawler!!
func getDirsCrawl() (func(string), <-chan string) {
	// gophers are goroutine counter!
	Gophers := int(1)
	DirsCache := make(map[string]bool)
	Mux := new(sync.Mutex)
	Ch := make(chan string, 10)

	var dirsCrawl func(string)
	dirsCrawl = func(dirname string) {
		defer func() {
			Mux.Lock()
			defer Mux.Unlock()
			Gophers--
			if Gophers == 0 {
				close(Ch)
			} else if Gophers < 0 {
				panic("crawl:Cannot managed Gophers!")
			}
		}()

		Mux.Lock()
		if DirsCache[dirname] {
			Mux.Unlock()
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
		for _, x := range info {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				Ch <- tmp
				Mux.Lock()
				Gophers++
				Mux.Unlock()
				go dirsCrawl(tmp)
			}
		}
	}
	return dirsCrawl, Ch
}
func processGetDirsCrawl() ([]string) {
	var dirs []string
	crawl, ch := getDirsCrawl()
	go crawl(*root)
	for x, ok := <-ch; ok; x, ok = <-ch {
		dirs = append(dirs, x)
	}
	return dirs
}

// Use wait group!!
func useWaitGroupCrawl(root string) (result []string) {
	// gophers are goroutine counter!
	DirsCache := make(map[string]bool)
	wg := new(sync.WaitGroup)
	mux := new(sync.Mutex)

	var dirsCrawl func(string)
	dirsCrawl = func(dirname string) {
		defer wg.Done()
		mux.Lock()
		if DirsCache[dirname] {
			return
		}
		DirsCache[dirname] = true
		mux.Unlock()

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
		for _, x := range info {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				mux.Lock()
				result = append(result, tmp)
				mux.Unlock()
				wg.Add(1)
				go dirsCrawl(tmp)
			}
		}
	}
	wg.Add(1)
	dirsCrawl(root)
	wg.Wait()
	return result
}
func processWaitGroupCrawl() []string {
	// direct return
	return useWaitGroupCrawl(*root)
}

func show(dirslist []string) {
	for _, x := range dirslist {
		fmt.Println(x)
	}
	fmt.Println("dirslist length", len(dirslist))
}

func equalCheck(T1, T2 []string) bool {
	var result bool
	fmt.Println("[]string length", len(T1), len(T2))
	sort.Strings(T1)
	sort.Strings(T2)

	// 素朴な実装
	// もう少しシンプルにしたい
	cache1 := make(map[string]bool)
	cache2 := make(map[string]bool)
	for _, x := range T1 {
		cache1[x] = true
	}
	for _, x := range T2 {
		cache2[x] = true
	}

	for _, x := range T1 {
		if !cache2[x] {
			log.Printf("T1 hove %v, but do not find in T2\n", x)
			result = false
		}
	}
	for _, x := range T2 {
		if !cache1[x] {
			log.Printf("T2 hove %v, but do not find in T1\n", x)
			result =false
		}
	}
	return result
}

func main() {
	fmt.Println(*root)
	fmt.Println(filepath.Dir(*root))

	dirs1, errlist1 := processDirsCrawl()
	show(dirs1)
	fmt.Println("processDirsCrawl:end")
	time.Sleep(time.Second)

	dirs2, errlist2 := processHayabusa()
	show(dirs2)
	fmt.Println("processHayabusa:end")
	time.Sleep(time.Second)

	dirs3, errlist3 := processSimpleCrawl()
	show(dirs3)
	fmt.Println("processSimpleCrawl:end")
	time.Sleep(time.Second)

	dirs4 := processGetDirsCrawl()
	show(dirs4)
	fmt.Println("processGetDirsCrawl:end")
	time.Sleep(time.Second)

	dirs5 := processWaitGroupCrawl()
	show(dirs5)
	fmt.Println("processWaitGroupCrawl:end")


	// simple crawl dirs 3 is expected
	fmt.Println("test 1")
	equalCheck(dirs3, dirs1)

	fmt.Println("test 2")
	equalCheck(dirs3, dirs2)

	fmt.Println("test 3")
	equalCheck(dirs3, dirs4)

	fmt.Println("test 4")
	equalCheck(dirs3, dirs5)

	sort.Strings(errlist1)
	sort.Strings(errlist2)
	sort.Strings(errlist3)
	fmt.Println("errlist1:", errlist1)
	fmt.Println("errlist2:", errlist2)
	fmt.Println("errlist3:", errlist3)
}

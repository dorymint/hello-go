// TODOGather
package main

// 少し考える simple simple simple
// 幾つか方法がある
// goroutine が軽量で数十万スレッド扱えるらしいので
// ガンガン使ってみる実装にしてみたい

// directory list を先に作って構造を取る
// ハードリンクとソフトリンクによる重複回避

// 各ディレクトリへ(各ファイルでもいい)gopherを投げ、ファイルを取って中身を検証してもらう

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	// flag strings
	depth = flag.Int("depth", 0, "serch depth, int")
	root  = flag.String("root", "", "serch directory root path, string")
	// default setting with init()
	target = flag.String("todo", "TODO:", "search target string")
)

func usage() {
	fmt.Printf(`Usage of %s:
	-depth=<search depth>
	-root=<search root path>
`, os.Args[0])
}

func fatalIF(str string, err error) {
	if err != nil {
		log.Fatalf("%s:%v\n", str, err)
	}
}

func init() {
	var err error
	*root, err = os.Getwd()
	fatalIF("init", err)

	// TEST:TODO: delete
	*root = "../hello-go"

	flag.Parse()
	*root, err = filepath.Abs(*root)
	fatalIF("init", err)
}

func pwd() (dir string) {
	dir, err := os.Getwd()
	fatalIF("pwd", err)
	fmt.Println(dir)
	return
}

func flagView() error {
	if len(flag.Args()) != 0 {
		fmt.Println("cmd = ", os.Args)
		fmt.Println("Unknown option")
		for _, x := range flag.Args() {
			fmt.Println(x)
		}
		fmt.Println("-----| Usage |-----")
		usage()
		/* TODO: error hundling */
		os.Exit(1)
		return errors.New("Unknown Flags")
		/* error hundling */
	}
	fmt.Println("-----| flag view |-----")
	fmt.Println("depth = ", *depth)
	fmt.Println("root  = ", *root)
	fmt.Println("------ --------- ------")
	return nil
}

// return FileInfo is os.FileInfo
//func structerGather() (os.FileInfo, error) {
//	return
//}

// TODOGather get file list and return TODOList
// TODO:
func TODOGather() error {
	dir, err := os.Open(*root)
	fatalIF("os.Open", err)
	defer func() { fatalIF("TODOGather", dir.Close()) }()

	// reaadはファイル内のreadポインタを動かしたままにする
	fmt.Println("info to dir")
	dirnames, err := dir.Readdirnames(0)
	fatalIF("Readdirnames", err)
	for i, x := range dirnames {
		fmt.Println(i, x)
	}

	// readポインタが動いているのでdirnamesが見つけられない
	fmt.Println("info to verbose")
	info, err := dir.Readdir(0)
	fatalIF("f.Readdir", err)
	for i, x := range info {
		fmt.Println(i, x.IsDir(), x.Name())
	}

	return errors.New("test")
}

// read file, file select are another function job
func fileInGather(filename string) (str []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("fileInGather:%q\n", err)
	}
	defer func() {
		if errclose := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "fileInGatherClose:%q\n", errclose)
		}
	}()

	// TODO: identify executable
	// That is another function job?
	for sc, i := bufio.NewScanner(f), uint(1); sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			//log.Println("fileInGather:", err)
			//return nil, nil
			return nil, fmt.Errorf("%q\n", err)
		}
		if index := strings.Index(sc.Text(), *target); index != -1 {
			str = append(str, fmt.Sprintf("L%v:%s", i, sc.Text()[index:]))
		}
	}
	return
}

// TODO: 取り敢えず...ディレクトリネームだけを集めてみる
// 作りがかなり怪しい...もっとシンプルにしたい...
func dirsGather(searchRoot string) []string {
	result := make([]string, 0, 1000)
	ch := make(chan string, 10)
	dirsmap := make(map[string]bool)
	gophers := uint(1)
	mux := new(sync.Mutex)
	//var errorlog []string

	var dirsCrawl func(string)
	dirsCrawl = func(dirname string) {
		// error logging
		defer func() {
			if err := recover(); err != nil {
				log.Printf("dirsCrawl:%v\n", err)
				//errorlog = append(errorlog, fmt.Sprintf("dirsCrawl:%v\n", err))
			}
		}()
		defer func() {
			mux.Lock()
			gophers--
			mux.Unlock()
		}()

		// open
		dir, err := os.Open(dirname)
		if err != nil {
			panic(err)
		}
		// close
		defer func() {
			if errclose := dir.Close(); errclose != nil {
				log.Printf("errclose:%v", errclose)
			}
		}()

		// TODO: identify executable, select and push text file
		// get dirnames
		dirnames, err := dir.Readdir(0)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len(dirnames) == 0 {
			return
		}
		// push dirnames
		for _, x := range dirnames {
			if x.IsDir() {
				tmp := filepath.Join(dirname, x.Name())
				mux.Lock()
				if dirsmap[tmp] {
					mux.Unlock()
					continue
				}
				dirsmap[tmp] = true
				gophers++
				mux.Unlock()

				// recur
				go dirsCrawl(tmp)
				fmt.Println(tmp)
				ch <- tmp
			}
		}
	}
	go dirsCrawl(searchRoot)
	// wait gophers!
	for gophers != 0 {
		time.Sleep(time.Nanosecond)
		result = append(result, <-ch)
	}
	fmt.Println("result = ", len(result), gophers)
	//fmt.Println(errorlog)
	return result
}

// 帯域でtypeとmutexを結びつけてサーチ関数でgopher走らせれば
// 非同期で入力を邪魔しないサーチが簡単にできそう

//func alldirs() (result []string) {
//	var dirsmap map[string]bool
//	aldi := func(dirname string) {}
//	aldi = func(dirname string) {
//		tmp, err := dirsGather(dirname)
//		fatalIF("aldi:", err)
//		if tmp == nil {
//			return
//		}
//		for _, x := range tmp {
//			// mux lock
//			if dirsmap[x] {
//				// mux unlock
//				continue
//			}
//			result = append(result, x)
//			dirsmap[x] = true
//			// mux unlock
//			// go aldi(x)
//			aldi(x)
//		}
//	}
//	aldi(*root)
//	return
//}

func main() {
	flagView()

	//TODOGather()

	// test
	log.Println("call to dirsGather")
	dirsGather(*root)
	//	test := dirsGather(*root)
	//	for _, x := range test {
	//		fmt.Println(x)
	//	}
}

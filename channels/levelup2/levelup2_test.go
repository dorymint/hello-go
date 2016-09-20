package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	*root, err = filepath.Abs("../../")
	if err != nil {
		log.Fatalf("TestMain:%v\n", err)
	}
	os.Exit(m.Run())
}


func makeTempDir(dirs []string, filemap map[string][]string) (root string) {
	// filemap[dirs[at]]return []filename
	root, err := ioutil.TempDir("", "crawl")
	fatalIF("tempdir:", err)

	for _, x := range dirs {
		dir := filepath.Join(root, x)
		err := os.Mkdir(dir, 0700)
		fatalIF("mkdir:", err)
		for _, y := range filemap[x] {
			ioutil.WriteFile(y, nil, 0700)
		}
	}
	return root
}


func TestUseWaitGroup(t *testing.T) {

	// TODO:create expected

//	dirlist, infomap := useWaitGroupCrawl(/* TODO:mock */)

	// TODO:fatal test
	/// t.Fatalf(msg)
}

func BenchmarkUseWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useWaitGroupCrawl(*root)
	}
}


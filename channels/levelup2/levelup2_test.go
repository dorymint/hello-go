package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

var (
	tmpRoot = ""
	tmpDirs = []string{
		"A",
		"A/aA",
		"A/aB",
		"B",
		"B/bA",
		"B/bA/baA",
	}
	tmpFilesMap = map[string][]string{
		tmpDirs[0]: {
			"A.txt",
		},
		tmpDirs[1]: {
			"AA.1.txt",
			"AA.2.txt",
		},
		tmpDirs[2]: {},
		tmpDirs[3]: {
			"AAA.1.txt",
			"AAA.2.txt",
			"AAA.3.txt",
		},
	}
)

func TestMain(m *testing.M) {
	// flag
	var err error
	*root, err = filepath.Abs("../../")
	if err != nil {
		log.Fatalf("TestMain:%v\n", err)
	}


	// make temp
	tmpRoot = makeTempDir(tmpDirs, tmpFilesMap)
	defer func() {
		if err := os.RemoveAll(tmpRoot); err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("tmproot is = %v\n", tmpRoot)

	os.Exit(m.Run())
}

// filemap is map[dirname][]files
func makeTempDir(dirs []string, filemap map[string][]string) (root string) {

	tmproot, err := ioutil.TempDir("", "crawl")
	fatalIF("tempdir", err)

	// create tempdir
	for _, x := range dirs {
		dirpath := filepath.Join(tmproot, x)
		err := os.MkdirAll(dirpath, 0700)
		fatalIF("mkdir", err)

		// create tempfile
		for _, y := range filemap[x] {
			filepath := filepath.Join(dirpath, y)
			err := ioutil.WriteFile(filepath, nil, 0700)
			fatalIF("ioutil.WriteFile", err)
		}
	}
	return tmproot
}

func deepEqualStrings(t *testing.T, expected, out []string) {
	if reflect.DeepEqual(expected, out) { return }
	t.Error("Not Equal!")
	t.Error("expected")
	for _, x := range expected {
		t.Error(x)
	}
	t.Error("but out")
	for _, x := range out {
		t.Error(x)
	}
	t.FailNow()
}

func TestUseWaitGroup(t *testing.T) {
	var expectedDirs []string
	for _, x := range tmpDirs {
		expectedDirs = append(expectedDirs, filepath.Join(tmpRoot, x))
	}

	// Run
	outDirs, outInfosMap := useWaitGroupCrawl(tmpRoot)

	// dirs check
	sort.Strings(outDirs)
	deepEqualStrings(t, expectedDirs, outDirs)

	// Create expected os.FileInfo Map
	expectedFileInfosMap := make(map[string][]os.FileInfo)
	for _, dirname := range append(expectedDirs, tmpRoot) {
		f, err := os.Open(dirname)
		if err != nil {
			t.Fatal(err)
		}
		info, err := f.Readdir(0)
		if err != nil {
			t.Fatal(err)
		}
		expectedFileInfosMap[dirname] = info
		if err := f.Close(); err != nil { t.Fatal(err) }
	}

	// CreateTestData expected info names
	var expectedNames []string
	for _, infos := range expectedFileInfosMap {
		for _, info := range infos{
			expectedNames = append(expectedNames, info.Name())
		}
	}
	sort.Strings(expectedNames)
	// CreateTestData out info names
	var outNames []string
	for _, infos := range outInfosMap {
		for _, info := range infos {
			outNames = append(outNames, info.Name())
		}
	}
	// Check []os.FileInfo names
	sort.Strings(outNames)
	deepEqualStrings(t, expectedNames, outNames)
}

func BenchmarkUseWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useWaitGroupCrawl(*root)
	}
}

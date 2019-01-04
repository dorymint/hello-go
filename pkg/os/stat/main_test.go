package main

import (
	"os"
	"path/filepath"
	"testing"
)

func makeRun(t *testing.T, testTargetFunc func(string) (os.FileInfo, error)) func(path string) {
	return func(path string) {
		t.Helper()
		f, err := testTargetFunc(path)
		if err != nil {
			t.Error(err)
		}
		t.Logf("[path]:%v\t[name]:%v\t[mode]:%v", path, f.Name(), f.Mode())
	}
}

var (
	txt = filepath.Join("testdata", "file.txt")
	ln  = filepath.Join("testdata", "file.ln")
)

func TestStat(t *testing.T) {
	run := makeRun(t, os.Stat)

	run(txt)
	run(ln)
}

func TestLstat(t *testing.T) {
	run := makeRun(t, os.Lstat)

	run(txt)
	run(ln)
}

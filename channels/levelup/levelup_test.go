package main

import (
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

func BenchmarkDirsCrawl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = processDirsCrawl()
	}
}


func BenchmarkHayabusaCrawl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = processHayabusa()
	}
}


func BenchmarkSimpleCrawl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = processSimpleCrawl()
	}
}

func BenchmarkGetDirsCrawl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = processGetDirsCrawl()
	}
}

func BenchmarkUseWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = processWaitGroupCrawl()
	}
}


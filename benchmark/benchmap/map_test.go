package benchmap

import (
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkMapBool(b *testing.B) {
	s := ""
	m := make(map[string]bool)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchMapBool(s, m)
	}
}
func BenchmarkMapBool_AddBool(b *testing.B) {
	s := ""
	m := make(map[string]bool)
	m["passive"] = true

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchMapBool(s, m)
	}
}

func BenchmarkMapInfo(b *testing.B) {
	s := ""
	m := make(map[string][]os.FileInfo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchMapInfo(s, m)
	}
}

func BenchmarkMapInfo_AddInfos(b *testing.B) {
	s := ""
	m := make(map[string][]os.FileInfo)
	pwd, err := os.Getwd()
	if err != nil {
		b.Fatal(err)
	}
	infos, err := ioutil.ReadDir(pwd)
	if err != nil {
		b.Fatal(err)
	}
	m["passive"] = infos

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchMapInfo(s, m)
	}
}

func BenchmarkMapBool_Add3(b *testing.B) {
	s := ""
	m := make(map[string]bool)
	m["test"] = true
	m["test2"] = true
	m["test3"] = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchMapBool(s, m)
	}
}
func BenchmarkMapInfo_AddInfos3(b *testing.B) {
	s := ""
	m := make(map[string][]os.FileInfo)
	pwd, err := os.Getwd()
	if err != nil {
		b.Fatal(err)
	}
	infos, err := ioutil.ReadDir(pwd)
	if err != nil {
		b.Fatal(err)
	}
	m["test"] = infos
	m["test1"] = infos
	m["test2"] = infos

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchMapInfo(s, m)
	}
}

// benchmark.
//
//	go test -bench .
//
package main

import (
	"testing"
)

var Map = func() map[string]*string {
	s := "world"
	return map[string]*string{
		"hello": &s,
	}
}()

// key val
var Pare = map[string]string{
	"hello": "world",
	"foo":   "bar",
}

func IF1(pare map[string]string) bool {
	nmap := len(Map)
	if nmap != 0 {
		for key, val := range pare {
			p, ok := Map[key]
			if ok && (p == nil || *p == val) {
				nmap--
			}
		}
	}
	return nmap == 0
}

func IF2(pare map[string]string) bool {
	nmap := len(Map)
	if nmap != 0 {
		for key, val := range pare {
			p, ok := Map[key]
			if !ok {
				continue
			}
			if p == nil || *p == val {
				nmap--
			}
		}
	}
	return nmap == 0
}

func BenchmarkIF1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IF1(Pare)
	}
}

func BenchmarkIF2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IF2(Pare)
	}
}

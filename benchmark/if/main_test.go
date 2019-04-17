// benchmark/if.
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
var Paire = map[string]string{
	"hello": "world",
	"foo":   "bar",
}

func If(paire map[string]string) bool {
	nmap := len(Map)
	if nmap != 0 {
		for key, val := range paire {
			p, ok := Map[key]
			if ok && (p == nil || *p == val) {
				nmap--
			}
		}
	}
	return nmap == 0
}

func IfContinue(paire map[string]string) bool {
	nmap := len(Map)
	if nmap != 0 {
		for key, val := range paire {
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

func BenchmarkIf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		If(Paire)
	}
}

func BenchmarkIfContinue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IfContinue(Paire)
	}
}

// benchmark.
//
//	go test -bench Shift
//
package main

import (
	"fmt"
	"testing"
)

var tests = []string{"hello", "world", "foo", "bar", "fizz", "buzz"}

func BenchmarkShiftAppend(b *testing.B) {
	ss := make([]string, 3)
	ps := fmt.Sprintf("%p", ss)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range tests {
			ss = append(ss[1:], s)
		}
	}
	b.StopTimer()
	b.Logf("before p:%s after p:%p", ps, ss)
}

func BenchmarkShiftDirect(b *testing.B) {
	ss := make([]string, 3)
	ps := fmt.Sprintf("%p", ss)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range tests {
			ss[0], ss[1], ss[2] = ss[1], ss[2], s
		}
	}
	b.StopTimer()
	b.Logf("before p:%s after p:%p", ps, ss)
}

func BenchmarkShiftFor(b *testing.B) {
	ss := make([]string, 3)
	index := len(ss) - 1
	ps := fmt.Sprintf("%p", ss)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range tests {
			for j := 0; j < index; j++ {
				ss[j] = ss[j+1]
			}
			ss[index] = s
		}
	}
	b.StopTimer()
	b.Logf("before p:%s after p:%p", ps, ss)
}

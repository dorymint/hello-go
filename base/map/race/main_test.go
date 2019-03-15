// race condition.
//
//	go test -v race
//
package main

import (
	"testing"
)

// read only
func TestNotRace(t *testing.T) {
	m := map[string]string{
		"hello": "world",
		"cat":   "dog",
	}
	go func() {
		t.Log(m["hello"])
	}()
	go func() {
		t.Log(m["cat"])
	}()
}

// read write
func TestRace(t *testing.T) {
	m := map[string]string{"hello": "world"}
	go func() {
		_ = m["hello"]
	}()
	go func() {
		m["fizz"] = "buzz"
	}()
}

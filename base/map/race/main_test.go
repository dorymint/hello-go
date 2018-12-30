// go test -v -race
package main

import (
	"testing"
)

func TestNotRace(t *testing.T) {
	m := map[string]string{
		"hello": "world",
		"cat":   "dog",
	}
	go func() {
		s := m["hello"]
		t.Log(1, s)
	}()
	go func() {
		s := m["hello"]
		t.Log(2, s)
	}()
	go func() {
		s := m["cat"]
		t.Log(3, s)
	}()
}

func TestIsRace(t *testing.T) {
	m := map[string]string{"hello": "world"}
	go func() {
		t.Log(m["hello"])
	}()
	go func() {
		m["cat"] = "dog"
	}()
}

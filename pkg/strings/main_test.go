package main

import (
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	str := "192.168.1.1:8080"
	t.Log("str:", str)
	t.Log("trim use index:", str[:strings.IndexRune(str, ':')])
}

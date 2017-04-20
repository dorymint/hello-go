package main

import (
	"reflect"
	"testing"
)

func TestGomemNew(t *testing.T) {
	g := gomemNew("title", "", nil)
	expected := &Gomem{Title: "title"}
	if !reflect.DeepEqual(g, expected) {
		t.Fatalf("Expected %q, but %q:", expected, g)
	}
}


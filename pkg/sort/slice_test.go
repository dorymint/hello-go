package sort_test

import (
	"sort"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []string{
		"hello",
		"wolrd",
		"a",
		"b",
	}
	t.Log("befor:", s)
	sort.Strings(s)
	t.Log("after:", s)
}

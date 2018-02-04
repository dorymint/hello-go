package sort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	ul := []uint{50, 2, 1, 9}
	t.Log("before:", ul)
	sort.Slice(ul, func(i, j int) bool {
		return ul[i] < ul[j]
	})
	t.Log("after:", ul)
}

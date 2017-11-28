package sort

import (
	"sort"
	"testing"
)

type title struct {
	name    string
	version []int
}

type titleList []title

func (tl titleList) Len() int {
	return len(tl)
}

func (tl titleList) Swap(i, j int) {
	tl[i], tl[j] = tl[j], tl[i]
}

func (tl titleList) Less(i, j int) bool {
	if len(tl[i].version) == len(tl[j].version) {
		for depth := range tl[i].version {
			if tl[i].version[depth] < tl[j].version[depth] {
				return true
			}
		}
	}
	return len(tl[i].version) < len(tl[j].version)
}

func TestSortTitleList(t *testing.T) {
	tests := titleList{
		{name: "02", version: []int{0, 1}},
		{name: "04", version: []int{0, 3}},
		{name: "03", version: []int{0, 2}},
		{name: "01", version: []int{0, 0}},

		{name: "00", version: []int{0}},
		{name: "-1", version: []int{}},
	}

	t.Run("sort.Sort", func(t *testing.T) {
		tmp := tests
		t.Log("before:", tmp)
		sort.Sort(tmp)
		t.Log("after:", tmp)
	})

	t.Run("sort.Slice", func(t *testing.T) {
		tmp := tests
		t.Log("before:", tmp)
		sort.Slice(tmp, func(i, j int) bool {
			if len(tmp[i].version) == len(tmp[j].version) {
				for depth := range tmp[i].version {
					if tmp[i].version[depth] < tmp[j].version[depth] {
						return true
					}
				}
			}
			return len(tmp[i].version) < len(tmp[j].version)
		})
		t.Log("after:", tmp)
	})
}

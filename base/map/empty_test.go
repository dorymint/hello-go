package map_test

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	m := make(map[string]struct{ str string })
	s := m["empty"].str
	t.Logf("%#v", s)
}

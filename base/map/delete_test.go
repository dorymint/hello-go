package map_test

import (
	"testing"
)

func TestBase(t *testing.T) {
	m := make(map[string]bool)
	m["hi"] = true
	t.Log(m)
	delete(m, "hi")
	t.Log(m)
	delete(m, "hi")
	t.Log(m)
}

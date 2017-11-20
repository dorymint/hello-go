package fp

import (
	"path/filepath"
	"testing"
)

func TestJoin(t *testing.T) {
	t.Log(filepath.Join("", "hello.txt"))
	t.Log(filepath.Join("", ".hello"))

	in := filepath.Join("", ".hello")
	out, err := filepath.Abs(in)
	t.Log("in:", in, "out:", out, "err:", err)
}

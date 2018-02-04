package fp

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestRel(t *testing.T) {
	base := "/base/root"
	path := []string{
		"file.txt",
		"../file.txt",
		"../../file.txt",
		"../../../file.txt",
		"dir/../../../another.txt",
		"../file.txt/..//tset.txt//",
	}

	t.Log("--- filepath.Join(base, filepath.Clean(path)) ---")
	for i, x := range path {
		path[i] = filepath.Join(base, filepath.Clean(x))
		t.Log("i:", i, "\npath:", x, "\nclean:", filepath.Clean(x), "\njoin:", path[i])
	}

	t.Log("--- filepath.Rel(base, path) ---")
	for i, x := range path {
		out, err := filepath.Rel(base, x)
		if err != nil {
			t.Log(err); continue
		}
		str := fmt.Sprintln(i, "base:", base)
		str += fmt.Sprintf("in:\t%#v\n", x)
		str += fmt.Sprintf("out:\t%#v\n", out)
		t.Log(str)
	}
}

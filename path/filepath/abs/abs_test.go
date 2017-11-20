package abs

import (
	"path/filepath"
	"testing"
)

func TestAbs(t *testing.T) {
	path := []string {
		"./", // pwd
		"./../",
		"../file",
		"", // pwd
		"/file",
		"TestAbs/file", // joined current working directory
		"/dir/../../../file",
	}
	for i, x := range path {
		p, err := filepath.Abs(x)
		if err != nil {
			t.Log(err)
			continue
		}
		t.Logf("%d\nin:\t%#v\nout:\t%#v\n", i, x, p)
	}
}

func TestJoin(t *testing.T) {
	in := filepath.Join("", ".hello")
	out, err := filepath.Abs(in)
	t.Log("in:", in, "out:", out, "err:", err)
}

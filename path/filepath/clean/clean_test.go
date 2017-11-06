package clean

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestClean(t *testing.T) {
	base := "/a"
	path := []string{
		"/a/b/c",
		"/a/../../b/c",
		"/a/b/../c",
		"../a/b/c",
	}
	t.Log("path:")
	for i, x := range path {
		str := fmt.Sprintln(i)
		str += fmt.Sprintln(x)
		str += fmt.Sprintln(filepath.Clean(x))
		str += fmt.Sprintln("appended base:", filepath.Join(base, filepath.Clean(x)))
		t.Log(str)
	}
}

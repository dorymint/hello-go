package p

import (
	"path"
	"testing"
)

func TestClean(t *testing.T) {
	unix := `/../test`
	t.Log("unix:", unix, "to", path.Clean(unix))

	win := `c:\windows\..\path`
	t.Log("windows:", win, "to", path.Clean(win))
}

func TestJoin(t *testing.T) {
	ss := []string{
		"hello",
		"world",
	}
	t.Log(path.Join(ss...))

	ss2 := []string{
		"http://hello///world/",
		"testbase.html",
	}
	t.Log(path.Join(ss2...))
}

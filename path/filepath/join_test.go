package fp

import (
	"path/filepath"
	"testing"
)

func TestJoin(t *testing.T) {
	t.Log(filepath.Join("", "hello.txt"))
	t.Log(filepath.Join("", ".hello"))
}

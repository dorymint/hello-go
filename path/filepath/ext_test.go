package fp

import (
	"path/filepath"
	"testing"
)

func TestExt(t *testing.T) {
	t.Log(filepath.Ext("base.txt"))
}

package slice_test

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var slice []string
	t.Logf("%#v", slice)
	slice = append(slice, "hello")
	t.Logf("%#v", slice)
}

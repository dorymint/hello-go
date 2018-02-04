package slice

import (
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	s1 := []string{"hello", "world"}
	t.Logf("%#v", reflect.DeepEqual(s1, s1))

	s2 := append(s1, "!!")
	t.Logf("%#v", reflect.DeepEqual(s1, s2))

	// NOTE: this is can't
	// s1 == s2

	// NOTE: this is accept
	// s1 == nil
}

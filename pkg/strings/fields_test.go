package strings_test

import (
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	s1 := "cmdname args1 args2"
	t.Log(s1,strings.Fields(s1))

	s2 := " cmdname  args1	args2		"
	t.Log(s2,strings.Fields(s2))
}

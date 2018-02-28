package strings_test

import (
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	s1 := "cmdname args1 args2"
	t.Logf("%s: %#v\n", s1, strings.Fields(s1))

	s2 := " cmdname  args1	args2		"
	t.Logf("%s: %#v\n", s2, strings.Fields(s2))

	s3 := "cmd \"space test\" arg"
	t.Logf("%s: %#v\n", s3, strings.Fields(s3))
}

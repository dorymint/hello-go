package regexp_test

import (
	"regexp"
	"testing"
)

func TestSimple(t *testing.T) {
	reg := "^[a-z]*"
	match := regexp.MustCompile(reg)
	str := "hello world"
	loc := match.FindStringIndex(str)
	if loc != nil {
		t.Log("loc:", loc, str[loc[0]:loc[1]])
	}
}

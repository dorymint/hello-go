package regexp_test

import (
	"regexp"
	"testing"
)

func TestSimple(t *testing.T) {
	reg := "^[a-z]+"
	match := regexp.MustCompile(reg)
	str := "hello world"
	loc := match.FindStringIndex(str)
	if loc != nil {
		t.Log("loc:", loc, str[loc[0]:loc[1]])
	}
}

// TODO: impl
func TestFind(t *testing.T) {
	reg := "[0-9]+"
	match := regexp.MustCompile(reg)
	str := "test 123456 test" + "\n" + "test" + "\n" + "7890 test"
	loc := match.FindStringIndex(str)
	t.Log(loc)
	t.Log(str[loc[0]:loc[1]])
}

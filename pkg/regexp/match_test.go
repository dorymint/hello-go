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

func TestFind(t *testing.T) {
	reg := "[0-9]+"
	match := regexp.MustCompile(reg)
	str := "test 123456 test" + "\n" + "test" + "\n" + "7890 test"
	loc := match.FindStringIndex(str)
	t.Log(loc)
	t.Log(str[loc[0]:loc[1]])
}

func TestGroup(t *testing.T) {
	reg := "^hello-(?:world|lily)$"
	match := regexp.MustCompile(reg)
	tests := []struct {
		str     string
		ismatch bool
	}{
		{str: "hello world", ismatch: false},
		{str: "hello-world", ismatch: true},
		{str: "hello-lily", ismatch: true},
		{str: "hello-", ismatch: false},
		{str: "hello-wo", ismatch: false},
		{str: "hello", ismatch: false},
		{str: "world", ismatch: false},
	}
	t.Logf("reg:%v", reg)
	for i, test := range tests {
		if b := match.MatchString(test.str); b != test.ismatch {
			t.Errorf("[Error %d]: unexpected result: reg:%v str:%v ismatch:%v bool:%v", i, reg, test.str, test.ismatch, b)
		} else if b {
			t.Logf("[Log %d]: matched:%v", i, test.str)
		}
	}
}

package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestFscan(t *testing.T) {
	tests := []struct {
		str     string
		exp     int
		wanterr bool
	}{
		{str: "1234", exp: 1234, wanterr: false},
		{str: " 1234", exp: 1234, wanterr: false},
		{str: "1234 ", exp: 1234, wanterr: false},
		{str: " 1234 ", exp: 1234, wanterr: false},
		{str: "1234 1", exp: 1234, wanterr: false},
		{str: "1 1234", exp: 1, wanterr: false},
	}

	var b []byte
	buf := bytes.NewBuffer(b)
	for i, test := range tests {
		buf.Reset()
		if _, err := buf.WriteString(test.str); err != nil {
			t.Fatal(err)
		}
		out, err := fscan(buf, ioutil.Discard)
		if err != nil {
			if test.wanterr {
				t.Logf("expected error: %v", err)
				continue
			}
			t.Fatal(err)
		}
		if test.exp != out {
			t.Errorf("FAILE: case %d:str=%#v\texp=%#v but out=%#v\n", i, test.str, test.exp, out)
		} else {
			t.Logf("case %d:str=%#v\texp=%v out=%d\n", i, test.str, test.exp, out)
		}
	}
}

func TestSscan(t *testing.T) {
	tests := []struct {
		str     string
		exp     int
		wanterr bool
	}{
		{str: "4321", exp: 4321, wanterr: false},
		{str: " 4321", exp: 4321, wanterr: false},
		{str: " 4321 ", exp: 4321, wanterr: false},
		{str: "4321 ", exp: 4321, wanterr: false},
		{str: "	4321	", exp: 4321, wanterr: false},
		{str: "4321 1", exp: 4321, wanterr: false},
		{str: "1 4321", exp: 1, wanterr: false},
	}

	for i, test := range tests {
		out, err := sscan(test.str)
		if err != nil {
			if test.wanterr {
				t.Logf("expected error: %v", err)
				continue
			}
			t.Fatal(err)
		}
		if out != test.exp {
			t.Errorf("FAIL case %d\texp=%#v but out=%#v\n", i, test.exp, out)
		} else {
			t.Logf("case %d\texp=%#v out=%#v\n", i, test.exp, out)
		}
	}
}

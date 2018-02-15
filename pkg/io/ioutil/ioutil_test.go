package ioutil_test

import (
	"io/ioutil"
	"testing"
)

func TestIoutil(t *testing.T) {
	write := func(fname string) {
		t.Log("case: ", fname)
		if err := ioutil.WriteFile(fname, []byte("hello"), 0600); err != nil {
			t.Log(err)
		}
		t.Log("writed")
	}
	write("t/dir")
	write("t/hello_1.txt")
	write("t/./hello_2.txt")
	write("t/nest/hello_3.txt")
	write("t/nest/notexist/hello_4.txt")
}

func TestReadDir(t *testing.T) {
	tests := []struct{
		in string
		err bool
	}{
		{ "t", false },
		{ "t/hello_1", true },
	}

	for i, test := range tests {
		t.Log(i, "in:", test.in)
		infos, err := ioutil.ReadDir(test.in)
		if err == nil && test.err {
			t.Error(i, "expected error but nil")
		} else {
			t.Log(i, "err:", err)
		}
		t.Log(i, "out infos:", infos)
		for _, info := range infos {
			t.Log(i, info.Name())
		}
	}
}

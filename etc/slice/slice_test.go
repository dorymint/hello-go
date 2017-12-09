package slice

import (
	"testing"
)

func TestSimple(t *testing.T) {
	f := func(s []string) {
		t.Log("[]string{}:", s)
		defer func() {
			if e := recover(); e != nil {
				t.Log(e, ": s[0]")
				return
			}
		}()
		t.Log("s[0]:", s[0])
		t.Log("s[1]:", s[1:])
		t.Log("s[2]:", s[2:])
		t.Log("exit")
	}

	t.Log("--- test []string{} ---")
	f([]string{})
	t.Log("--- test []string{\"hi\"} ---")
	f([]string{"hi"})
}

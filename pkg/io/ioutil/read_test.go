package ioutil_test

import (
	"io/ioutil"
	"testing"
)

func TestRead(t *testing.T) {
	b, err := ioutil.ReadFile("")
	t.Log(b, err)
}

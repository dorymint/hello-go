package md5_test

import (
	"crypto/md5"
	"crypto/sha512"
	"io"
	"testing"
)

func TestSum(t *testing.T) {
	md5w := md5.New()

	hello := "hello"
	t.Log(hello, md5.Sum([]byte(hello)))
	io.WriteString(md5w, hello)
	t.Log(hello, md5w.Sum(nil))

	world := "world"
	t.Log(world, md5.Sum([]byte(world)))
	io.WriteString(md5w, world)
	t.Log(world, md5w.Sum(nil), "is appended hello")

	md5w.Reset()
	io.WriteString(md5w, world)
	t.Log(world, md5w.Sum(nil), "after reset then append world")
}

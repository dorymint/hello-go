package sha512_test

import (
	"crypto/sha512"
	"io"
	"testing"
)

func TestCheck(t *testing.T) {
	sha := sha512.New512_256()

	t.Log("nil", sha.Sum(nil))

	hello := "hello"
	io.WriteString(sha, hello)
	t.Log(hello, sha.Sum(nil))
}

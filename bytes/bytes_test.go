package tbytes

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	var b bytes.Buffer
	t.Run("write", func(t *testing.T) {
		b.Write([]byte("hello"))
		t.Log("string:", b.String(), "bytes:", b.Bytes())
	})

	t.Run("count", func(t *testing.T) {
		echo := func() func() []byte {
			b := []byte("lily\n")
			return func() []byte {
				b = append(b, b...)
				return b
			}
		}()
		t.Log("push \"lily\\n\"")
		for i := 0; i < 10; i++ {
			b.Write(echo())
		}
		t.Log("count of [lily]", bytes.Count(b.Bytes(), []byte("lily")))
		b.Reset()
		t.Log("after reset", b.String(), b.Bytes())
	})
}

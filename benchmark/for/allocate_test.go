package allocate

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func newErr(i int) error {
	return fmt.Errorf("%d", i)
}

func equal() {
	for i := 0; i < 10; i++ {
		if err := newErr(i); err != nil {
			fmt.Fprintf(ioutil.Discard, "%v %p %p\n", err, &err, err)
		}
	}
}

func colonEqual() {
	var err error
	for i := 0; i < 10; i++ {
		if err = newErr(i); err != nil {
			fmt.Fprintf(ioutil.Discard, "%v %p %p\n", err, &err, err)
		}
	}
}

func BenchmarkEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		equal()
	}
}

func BenchmarkColonEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		colonEqual()
	}
}

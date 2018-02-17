package interface_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type One struct {
	w    io.Writer
	name string
}

func (one *One) Sub(args []string) error {
	_, err := fmt.Fprintln(one.w, one.name, args)
	return err
}

type Two string

func (two Two) Sub(args []string) error {
	return fmt.Errorf("%v %v", two, args)
}

type I interface {
	Sub(args []string) error
}

func TestInterface(t *testing.T) {
	var i I
	buf := bytes.NewBufferString("")

	i = &One{
		w:    buf,
		name: "structure of One",
	}
	err := i.Sub([]string{"One", "hello", "world"})
	t.Log("err:", err)
	t.Log("buf:", buf)

	i = Two("structer of Two")
	err = i.Sub([]string{"Two", "hello", "world"})
	t.Log("err:", err)
}

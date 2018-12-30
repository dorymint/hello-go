package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"
)

// TODO: impl

func TestInteractive(t *testing.T) {
	cmdline := `hello world
exit
`
	wbuf := bytes.NewBufferString("")
	rbuf := bytes.NewBufferString(cmdline)

	http.HandleFunc("/", rootHandler)
	interactive(wbuf, rbuf)
	go t.Log(http.ListenAndServe("localhost:8080", nil))
	time.Sleep(time.Second * 10)
	t.Log("wbuf:", wbuf)
}

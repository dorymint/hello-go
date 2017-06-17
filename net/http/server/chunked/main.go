package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var image struct {
	file *os.File
	b    []byte
	mux  sync.Mutex
}

func updateImage() error {
	image.mux.Lock()
	defer image.mux.Unlock()
	var err error
	if image.b, err = ioutil.ReadAll(image.file); err != nil {
		return err
	}
	if _, err = image.file.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

func getb() []byte {
	image.mux.Lock()
	defer image.mux.Unlock()
	b := image.b
	return b
}

func ImageUpdateServer(w http.ResponseWriter, req *http.Request) {
	log.Printf("connect:%+v\t%+v\n", req.RemoteAddr, req.RequestURI)
	var b []byte
	buf := bytes.NewBuffer(b)

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("flusher not impl")
	}

	for {
		buf.Write(getb())
		w.Write(buf.Bytes())
		buf.Reset()
		flusher.Flush()
		time.Sleep(time.Second)
	}
}

func init() {
	f, err := os.Open("plan9.png")
	if err != nil {
		panic(err)
	}
	// do not close
	image.file = f
	if err := updateImage(); err != nil {
		panic(err)
	}
}
func main() {
	go func() {
		var err error
		if err = updateImage(); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 5)
	}()
	http.HandleFunc("/", ImageUpdateServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

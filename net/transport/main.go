package main

import (
	"fmt"
	"net/http"
)

func main() {
	show := func(tr *http.Transport) {
		fmt.Println("----- check default transport -----")
		fmt.Println("Compress:", tr.DisableCompression)
		fmt.Println("KeepAlive:", tr.DisableKeepAlives)
		fmt.Println("ExpectContinueTimeout:", tr.ExpectContinueTimeout)
		fmt.Println("IdleConnTimeout::", tr.IdleConnTimeout)
		fmt.Println("MaxIdleConns:", tr.MaxIdleConns)
		fmt.Println("MaxIdleConnsPerHost:", tr.MaxIdleConnsPerHost)
		fmt.Println("TLSConfig:", tr.TLSClientConfig)
		fmt.Println("TLSHandshakeTimeout:", tr.TLSHandshakeTimeout)
		fmt.Println("TLSNextProto,map:", tr.TLSNextProto)
	}
	tr := &http.Transport{}
	show(tr)
	/// TODO:
}

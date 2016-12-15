
// web cliant

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", "http://localhost:6060", "")

func init(){
	flag.Parse()
}

func main() {
	fmt.Println("hello web client")
	// Use `godoc -http :6060`
	url := *addr //"http://localhost:6060"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()


	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArray))

	time.Sleep(time.Millisecond)
	fmt.Println("resp")
	fmt.Println("content length", resp.ContentLength)
	// NOTE: chunk と共存しない

	fmt.Println("header", resp.Header)
	// TODO: Research

	fmt.Println("proto", resp.Proto) // protocol
	fmt.Println("proto Major", resp.ProtoMajor)
	fmt.Println("proto Minor", resp.ProtoMinor)

	fmt.Println("proto request", resp.Request)
	// TODO: Research

	fmt.Println("status", resp.Status) // string
	fmt.Println("status code", resp.StatusCode) // int

	fmt.Println("TLS", resp.TLS) // TLS connection state
	// TODO: Research

	fmt.Println("Trailer", resp.Trailer) // http header
	// TODO: Research, what difference resp.Header
	// NOTE: Request body の後に追加で送られるヘッダー情報らしい?

	fmt.Println("Transfer Encoding", resp.TransferEncoding)
	// NOTE: chunked: content lentgth と共存しない

	fmt.Println("Uncompressed", resp.Uncompressed)
	// compress は圧縮の動詞, compressese が圧縮

	fmt.Println("cookies", resp.Cookies())
}


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
//var scheme = flag.String("scheme", "http:", "")
//var host = flag.String("host", "//localhost:6060", "")

func init(){
	flag.Parse()
}

func main() {
	fmt.Println("hello web client")
	time.Sleep(time.Second)
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
	time.Sleep(time.Second)
	fmt.Println(string(byteArray))
	time.Sleep(time.Second)

	fmt.Println("--- FOR DEBUG BEGIN ---")
	fmt.Println("--- RESP ---\n", resp)
	time.Sleep(time.Second)
	fmt.Println("--- CONTENT LENGTH ---\n", resp.ContentLength)
	time.Sleep(time.Second)
	// NOTE: chunk と共存しない

	fmt.Println("--- HEADER ---\n", resp.Header)
	time.Sleep(time.Second)

	fmt.Println("--- PROTO ---\n", resp.Proto) // protocol
	time.Sleep(time.Second)
	fmt.Println("--- PROTO MAJOR ---\n", resp.ProtoMajor)
	time.Sleep(time.Second)
	fmt.Println("--- PROTO MINOR ---\n", resp.ProtoMinor)
	time.Sleep(time.Second)

	fmt.Println("--- PROTO REQUEST ---\n", resp.Request)
	time.Sleep(time.Second)

	fmt.Println("--- STATUS ---\n", resp.Status) // string
	time.Sleep(time.Second)
	fmt.Println("--- STATUS CODE ---\n", resp.StatusCode) // int
	time.Sleep(time.Second)

	fmt.Println("--- TLS ---\n", resp.TLS) // TLS connection state
	time.Sleep(time.Second)

	fmt.Println("--- TRAILER ---\n", resp.Trailer) // http header
	time.Sleep(time.Second)
	// NOTE: Request body の後に追加で送られるヘッダー情報らしい?

	fmt.Println("--- TRANSFER ENCODING ---\n", resp.TransferEncoding)
	time.Sleep(time.Second)
	// NOTE: chunked: content lentgth と共存しない

	fmt.Println("--- UNCOMPRESSED ---\n", resp.Uncompressed)
	time.Sleep(time.Second)
	// compress は圧縮の動詞, compressese が圧縮

	fmt.Println("--- COOKIES ---\n", resp.Cookies())
	time.Sleep(time.Second)
	fmt.Println("--- FOR DEBUG END ---")
}

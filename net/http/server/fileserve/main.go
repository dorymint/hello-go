package main

import (
	"flag"
	"log"
	"net/http"
)

func logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.RequestURI)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	const root = "./doc"

	port := flag.String("port", "localhost:8080", "listen port")
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("invalid argument:", flag.Args())
	}

	http.Handle("/", logging(http.FileServer(http.Dir(root))))
	log.Fatal(http.ListenAndServe(*port, nil))
}

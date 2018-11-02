// References https://golang.org/doc/effective_go.html#control-structures

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Simple counter server
type CounterStruct struct {
	n int
}

func (c *CounterStruct) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	c.n++
	fmt.Fprintf(w, "counter=%d\n", c.n)
}

// Simple counter server
type CounterInt int

func (ci *CounterInt) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	*ci++
	fmt.Fprintf(w, "counter(int)=%d\n", *ci)
}

// Notification
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch <- r
	fmt.Fprintln(w, "notification sent")
}

func main() {
	cs := new(CounterStruct)
	http.Handle("/counter", cs)

	ci := new(CounterInt)
	http.Handle("/counter_int", ci)

	ch := make(Chan)
	http.Handle("/notify", ch)
	go func() {
		for {
			r := <-ch
			fmt.Println("information:", r)
		}
	}()

	// help
	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		h := []string{"/help", "/counter", "/counter_int", "/notify"}
		fmt.Fprintf(w, "%s\n", strings.Join(h, "\n"))
	})

	// Server Running
	log.Fatal(http.ListenAndServe(":8080", nil))
}

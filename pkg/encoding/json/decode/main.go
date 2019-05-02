// pkg/encoding/json/decode.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

var jsonStreams = []string{
	// object
	`{}`,

	// array
	`[]`,

	// string
	`"str"`,

	// number
	`0`,

	// value
	`null`,

	`{"hello":"world","foo":["bar","fizz","buzz"]}`,
}

func decode(str string) error {
	dec := json.NewDecoder(strings.NewReader(str))
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Printf("  %v\n", reflect.TypeOf(t))
		fmt.Printf("    %+v more:%v\n", t, dec.More())
	}
}

func run() error {
	for _, stream := range jsonStreams {
		fmt.Printf("stream: %q\n", stream)
		err := decode(stream)
		if err != nil {
			return err
		}
		fmt.Println()
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

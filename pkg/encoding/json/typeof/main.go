// pkg/encodig/json/typeof.
package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	"math"
	"reflect"
	"strconv"
)

var jsonStreams = []string{
	`0`,
	`null`,
	`true`,
	`"string"`,
	`{}`,
	`{"key":"val"}`,
	`{"key":null}`,
	`[]`,
	`[0]`,
	`["string"]`,
	`[0,"string"]`,
	`[[]]`,
	`["hello", ["world"]]`,
	`["hello", {}]`,

	`{
	"name":"foo",
	"age": 14,
	"items": {
		"p1": 1,
		"p2": "string"
	},
	"stock":[
		1,
		2,
		3
	],
	"live": true
}`,

	//	invalid
	`{{}}`,
	`[{{}}]`,

	// expected: 18446744073709551615
	// but out : 1.8446744073709552e+19
	strconv.FormatUint(math.MaxUint64, 10),
}

func main() {
	for _, test := range jsonStreams {
		fmt.Printf("json:%q\n", test)
		var v interface{}
		err := json.Unmarshal([]byte(test), &v)
		if err != nil {
			fmt.Printf("Err:%v\n\n", err)
			continue
		}

		t := reflect.TypeOf(v)
		if t == nil {
			continue
		}
		fmt.Printf("out:%v\n", t)
		if err := ast.Print(token.NewFileSet(), v); err != nil {
			panic(err)
		}
		fmt.Println()
	}
}

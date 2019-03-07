package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ref: golang.org/pkg/encoding/json

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}
	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}
	return json.Marshal(s)
}

func main() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		panic(err)
	}

	fmt.Printf("zoo:%q\n", zoo)
	fmt.Println("Tof zoo:", reflect.TypeOf(zoo))
	fmt.Println("Vof zoo:", reflect.ValueOf(zoo))
	// iota: gopher=1, unknown=0, zebra=2

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}
	fmt.Println("Gophers:", census[Gopher], "Zebra:", census[Zebra], "Unknown:", census[Unknown])

	// ref: www.kaoriya.net/blog/2016/06/25
	var v interface{}
	if err := json.Unmarshal([]byte(blob), &v); err != nil {
		panic(err)
	}
	fmt.Println("var v interface{}:", v)
	fmt.Printf("%#v\n", v)
	vv, ok := v.([]interface{})
	if !ok {
		panic("invalid")
	}
	fmt.Println("Tof v:", reflect.TypeOf(&v))
	fmt.Println("Tof vv:", reflect.TypeOf(&vv))
	vv = append(vv, []interface{}{32, 22}...)
	for _, v := range vv {
		switch v.(type) {
		case string:
			fmt.Printf("%v: %#v\n", reflect.TypeOf(v), v)
		default:
			fmt.Printf("not string: %g\n", v)
		}
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func helloJSON() {
	// Marshal
	b, err := json.Marshal(struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "lily",
		Age:  11,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// Unmarshal
	v := new(struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Nyan string `json:"nyan"`
	})
	fmt.Println(string(b[:len(b)-1]))
	err = json.Unmarshal(append(b[:len(b)-1], []byte(`,"nyan":"nyani"}`)...), v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nyan:", v.Nyan)
	fmt.Println("v:", v)
}

func compact() {
	b := new(bytes.Buffer)
	err := json.Compact(b, []byte(`[  {"name":"lily","age":11}  ,  {"name":"dory","age":12}  ]`))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

// FROM: golang.org/pkg/encoding/json
//     : custommarshal
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
func customMarshal() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
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

	// REF: www.kaoriya.net/blog/2016/06/25
	var v interface{}
	if err := json.Unmarshal([]byte(blob), &v); err != nil {
		log.Fatal(err)
	}
	fmt.Println("var v interface{}:", v)
	fmt.Printf("%#v\n", v)
	vv, ok := v.([]interface{})
	if !ok {
		log.Fatal("invalid")
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

func readJSON() {
	// JSONのkeyが大文字小文字の違いだけならstructにタグを書かなくてもマッピングしてくれる
	type taskJSON struct {
		Taskname    string
		Description string
		Content     string
	}
	var t []taskJSON
	b, err := ioutil.ReadFile("./readJSON.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b, &t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	for i, v := range t {
		fmt.Println(i)
		fmt.Println("  taskname:", v.Taskname)
		fmt.Println("  description:", v.Description)
		fmt.Println("  content:", v.Content)
	}
}

func writeJSON() {
	// 書き込むときはタグを付けないとそのまま大文字で出力されてしまう
	type taskJSON struct {
		Taskname    string `json:"taskname"`
		Description string `json:"description"`
		Content     string `json:"content"`
	}
	var t []taskJSON
	b, err := ioutil.ReadFile("./readJSON.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		log.Fatal(err)
	}
	t = append(t, taskJSON{
		Taskname:    "writeJSON",
		Description: "wirte test",
		Content:     "nyan",
	})

	// no indent
	out, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

	// indent
	out, err = json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

	if err := ioutil.WriteFile("./writeJSON.json", out, 0600); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	split("helloJSON")
	helloJSON()
	split("compact")
	compact()
	split("customMarshal")
	customMarshal()
	split("readJSON")
	readJSON()
	split("writeJSON")
	writeJSON()
}

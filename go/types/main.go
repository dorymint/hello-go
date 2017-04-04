package main

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"reflect"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func eval(expr string) types.TypeAndValue {
	tv, err := types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, expr)
	if err != nil {
		log.Fatal(err)
	}
	return tv
}

func typeInfo() {
	tvof := func(tv types.TypeAndValue) {
		fmt.Println(tv)
		fmt.Println("", tv.Type)
		fmt.Println("", tv.Value)
		fmt.Println("", reflect.ValueOf(tv.Value).Type())
	}
	tvof(eval("12 + 1"))
	tvof(eval("int(12) + int(1)"))
	tvof(eval("12.0 + 1.0"))
	tvof(eval("float64(12) + float64(1)"))
}

func main() {
	split("typeInfo")
	typeInfo()
}

package main

// Ref: http://qiita.com/tenntenn/items/f029425a844687a0e64b

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func parse() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./testdata/main.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Imports:")
	for _, i := range f.Imports {
		fmt.Printf("\t%+v\n", i)
	}

	fmt.Println("Comments:")
	for _, c := range f.Comments {
		fmt.Printf("\t%+v\n", c)
	}

	fmt.Println("Unresolved:")
	for _, i := range f.Unresolved {
		fmt.Printf("\t%+v\n", i)
	}

	fmt.Println("Decls:")
	for i := range f.Decls {
		fmt.Printf("\t%+v\n", f.Decls[i])
		fmt.Printf("\t%+v\n", f.Decls[i].Pos())
		fmt.Printf("\t%+v\n", fset.Position(f.Decls[i].Pos()))
	}

	fmt.Println("Position:")
	// at package
	fmt.Printf("\t%+v\n", fset.Position(f.Pos()))
	// at }
	fmt.Printf("\t%+v\n", fset.Position(f.End()))

	fmt.Println("Scope:")
	fmt.Printf("\t%+v\n", f.Scope)

	fmt.Println("Scope main:")
	fmt.Printf("\t%+v\n", f.Scope.Lookup("main"))
	fmt.Printf("\t%+v\n", f.Scope.Objects["main"].Kind)
	fmt.Printf("\t%+v\n", f.Scope.Objects["main"].Name)
	fmt.Printf("\t%+v\n", fset.Position(f.Scope.Objects["main"].Pos()))

	fmt.Println("Scope Lookup:")
	fmt.Printf("\t%+v\n", f.Scope.Lookup("hi"))
	fmt.Printf("\t%+v\n", f.Scope.Lookup("nyan"))
}

func expr() {
	e, err := parser.ParseExpr(`1+12`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\t%+v\n", e)
	fmt.Printf("\t%+v\n", e.Pos())
	fmt.Printf("\t%+v\n", e.End())

	nest := 1
	ast.Inspect(e, func(n ast.Node) bool {
		fmt.Printf("%s%[2]T %[2]v\n", strings.Repeat("\t", nest), n)
		if n != nil {
			nest++
		} else {
			nest--
		}
		return true
	})

	os.Stdout.WriteString("\t")
	format.Node(os.Stdout, token.NewFileSet(), e)
}

func main() {
	fmt.Println("parse():")
	parse()
	fmt.Printf("\n\n")

	fmt.Println("expr():")
	expr()
	fmt.Printf("\n\n")
}

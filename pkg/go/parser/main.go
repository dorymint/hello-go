package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func helloParser() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./t/nyan.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("%q\n", f)

	split("Imports")
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
	split("Comments")
	for _, s := range f.Comments {
		fmt.Println(s.Text())
	}
	split("Unresolved")
	for _, s := range f.Unresolved {
		fmt.Println(s)
	}
	split("Decls")
	for i := 0; i < len(f.Decls); i++ {
		fmt.Println(f.Decls[i])
		fmt.Println(f.Decls[i].Pos())
		fmt.Println(fset.Position(f.Decls[i].Pos()))
	}
	split("nyan.go position")
	// at package
	fmt.Println(fset.Position(f.Pos()))
	// at }
	fmt.Println(fset.Position(f.End()))

	split("Scope")
	fmt.Println(f.Scope)
	split("Scope main")
	fmt.Println(f.Scope.Lookup("main"))
	fmt.Println(f.Scope.Objects["main"].Kind)
	fmt.Println(f.Scope.Objects["main"].Name)
	fmt.Println(fset.Position(f.Scope.Objects["main"].Pos()))
	split("Scope Lookup")
	fmt.Println(f.Scope.Lookup("hi"))
	fmt.Println(f.Scope.Lookup("nyan"))
}

func expr() {
	a, err := parser.ParseExpr(`1+12`)
	if err != nil {
		panic(err)
	}
	fmt.Println(a.Pos())
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", a.Pos())
	fmt.Printf("%q\n", a.End())

	// FROM: http://qiita.com/tenntenn/items/f029425a844687a0e64b
	var i int
	ast.Inspect(a, func(n ast.Node) bool {
		fmt.Printf("%s%[2]T %[2]v\n", strings.Repeat(" ", i), n)
		if n != nil {
			i++
		} else {
			i--
		}
		return true
	})
	format.Node(os.Stdout, token.NewFileSet(), a)
	println()
}

func applay() {
	a, err := parser.ParseExpr(`1+12`)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func main() {
	split("helloParser()")
	helloParser()
	split("expr()")
	expr()
	split("applay()")
	applay()
}

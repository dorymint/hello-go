package main

import (
	"fmt"
	"go/parser"
	"go/token"
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

func main() {
	helloParser()
}

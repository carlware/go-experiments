package egs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	switch val := n.(type) {
	case *ast.StructType:
		// var vv visitor
		fmt.Println(strings.Repeat("\t", int(v)), val)
		// vv.Visit(val)
	}
	return v + 1
}

func TestVisitor() {
	file := "./models/user.go"

	// src := bytes.NewReader([]byte(""))
	_, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fs := token.NewFileSet()
	fil, err := parser.ParseFile(fs, file, nil, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(fil.Name.String())

	// for _, s := range fil.Decls {
	// 	fmt.Println(s)
	// }

	var v visitor
	ast.Walk(v, fil)

	// ast.Inspect(fil, func(n ast.Node) bool {
	// 	switch x := n.(type) {

	// 	case *ast.TypeSpec:
	// 		fmt.Printf("decl %s\n", x)
	// 	case *ast.StructType:
	// 		// fmt.Println("s", x.Fields.NumFields())
	// 		// for _, f := range x.Fields.List {
	// 		// 	fmt.Println("l", f.Type.Pos(), f.Names)
	// 		// 	fmt.Println("p")
	// 		// 	// for _, id := range f.Names {
	// 		// 	// fmt.Println("id", id.Name)
	// 		// 	// fmt.Printf("%s : %s \n", fs.Position(f.Pos()), id)
	// 		// 	// e := fs.Position(f.Pos())
	// 		// 	// id.
	// 		// 	// }
	// 		// }
	// 		// case *ast.ArrayType:
	// 		// 	fmt.Println("arr", x)
	// 	}
	// 	return true
	// })

	// ast.Print(fs, fil)

}

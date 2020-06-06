package egs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func TestStructInfo() {
	src := `
    package foo

    type Thing struct {
    Field1 string
    Field2 []int
    Field3 map[byte]float64
  }`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)

	if err != nil {
		panic(err)
	}

	// hard coding looking these up
	typeDecl := f.Decls[0].(*ast.GenDecl)
	structDecl := typeDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	fields := structDecl.Fields.List

	for _, field := range fields {
		typeExpr := field.Type

		start := typeExpr.Pos() - 1
		end := typeExpr.End() - 1

		// grab it in source
		typeInSource := src[start:end]

		fmt.Println(typeInSource)
	}

}

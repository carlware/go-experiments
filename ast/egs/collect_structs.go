package egs

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

const CODE = `
package models

type Foo struct {
  Name string
  Age  int
}
`

type structType struct {
	name string
	node *ast.StructType
}

type config struct {
	structName string
	fset       *token.FileSet
}

func (c *config) structSelection(file ast.Node) (int, int, error) {
	structs := collectStructs(file)

	var encStruct *ast.StructType
	for _, st := range structs {
		if st.name == c.structName {
			encStruct = st.node
		}
	}

	if encStruct == nil {
		return 0, 0, errors.New("struct name does not exist")
	}

	// struct selects all lines inside a struct
	start := c.fset.Position(encStruct.Pos()).Line
	end := c.fset.Position(encStruct.End()).Line

	return start, end, nil
}

// collectStructs collects and maps structType nodes to their positions
func collectStructs(node ast.Node) map[token.Pos]*structType {
	structs := make(map[token.Pos]*structType, 0)
	collectStructs := func(n ast.Node) bool {
		t, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if t.Type == nil {
			return true
		}

		structName := t.Name.Name

		x, ok := t.Type.(*ast.StructType)
		if !ok {
			return true
		}

		structs[x.Pos()] = &structType{
			name: structName,
			node: x,
		}
		return true
	}
	ast.Inspect(node, collectStructs)
	return structs
}

func getAst(code string, token *token.FileSet) *ast.File {
	f, err := parser.ParseFile(token, "example.go", code, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	return f
}

func TestCollectStructs() {
	fs := token.NewFileSet()
	f := getAst(CODE, fs)

	c := &config{
		structName: "Foo",
		fset:       fs,
	}

	s, e, err := c.structSelection(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("start", s, "end", e)

}

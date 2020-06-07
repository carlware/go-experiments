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

type Bar struct {
  Baz  Float
  Game bool
}

type Foo struct {
  Name string
  Age  int
  Baz  *Baz
}
`

type structType struct {
	name string
	node *ast.StructType
}

type config struct {
	structName string
	fset       *token.FileSet
	code       string
}

type Field struct {
	Name string
	Type string
	List []Field
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

func (c *config) getStruct(node ast.Node, start, end int) []Field {
	fields := []Field{}
	rewriteFunc := func(n ast.Node) bool {
		x, ok := n.(*ast.StructType)
		if !ok {
			return true
		}

		for _, f := range x.Fields.List {
			line := c.fset.Position(f.Pos()).Line

			if !(start <= line && line <= end) {
				continue
			}

			typeExpr := f.Type
			s := typeExpr.Pos() - 1
			e := typeExpr.End() - 1

			// grab it in source
			fieldType := c.code[s:e]

			fieldName := ""
			if len(f.Names) != 0 {
				for _, field := range f.Names {
					fieldName = field.Name
					break
				}
			}
			fields = append(fields, Field{fieldName, fieldType, nil})

			// nothing to process, continue with next line
			if fieldName == "" {
				continue
			}

			fmt.Println("field: ", fieldName, fieldType)
		}

		return true
	}

	ast.Inspect(node, rewriteFunc)
	return fields
}

func getFieldsByName(structs map[token.Pos]*structType, name string) []Field {
	fields := []Field{}

	var encStruct *ast.StructType
	for _, st := range structs {
		if st.name == c.structName {
			encStruct = st.node
		}
	}

	return fields
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

	structs := collectStructs(f)
	fmt.Println(structs)

	// c := &config{
	// 	structName: "Foo",
	// 	fset:       fs,
	// 	code:       CODE,
	// }
	// s, e, err := c.structSelection(f)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("start", s, "end", e)
	// st := c.getStruct(f, s, e)
	// fmt.Println(st)

	// buf := &bytes.Buffer{}
	// err = format.Node(buf, fs, f)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("//code// \n%s \n", buf.Bytes())
}

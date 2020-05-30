package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Foo string
}

func (a *A) PrintFoo() {
	fmt.Println("Foo value is " + a.Foo)
}

func TypeOf() {
	a := &A{Foo: "afoo"}
	val := reflect.Indirect(reflect.ValueOf(a))
	fmt.Println(val.Type().Field(0).Name)
}

func main() {

	TypeOf()
	// reflex.StructTypes()
	// reflex.CheckExportedMethodsAndFields()
}

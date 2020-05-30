package reflex

import (
	"fmt"
	"io"
	"reflect"
)

type User struct {
	ID string
}

type My struct {
	A int
	B interface{}
	C io.Reader
	U User
}

func (m *My) He(s string) {

}

func StructTypes() {
	t := reflect.TypeOf(My{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("Field %q, type: %-12v, type name: %-8q, is interface{}: %v\n",
			f.Name, f.Type,
			f.Type.Name(),
			f.Type.String() == "interface {}",
		)
	}
}

package eg

import "fmt"

type Stringer interface {
	String() string
}

type Struct1 struct {
	field1 string
}

func (s Struct1) String() string {
	return s.field1
}

type Struct2 struct {
	field1 []string
	dummy  bool
}

func (s Struct2) String() string {
	return fmt.Sprintf("%v, %v", s.field1, s.dummy)
}

type StringerContainer struct {
	Field string
	Stringer
}

func TestInterfaceContainer() {
	fmt.Println(StringerContainer{
		Field:    "Struct1",
		Stringer: Struct1{"This is Struct1"},
	})
	fmt.Println(StringerContainer{"Struct2", Struct2{[]string{"This", "is", "Struct1"}, true}})
}

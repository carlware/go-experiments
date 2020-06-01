package eg

import "fmt"

type User struct {
	Name    string
	Surname string
	Age     uint32
}

type Student struct {
	User
	Grade uint8
}

func (s Student) String() string {
	return fmt.Sprintf("I'am %s %d years old and I'am in the %d grade.", s.Name, s.Age, s.Grade)
}

func TestEmbedded() {
	st1 := Student{User{"c", "r", 20}, 6}
	fmt.Printf("%s", st1)
}

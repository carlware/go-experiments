package main

import (
	"carlware/hasfield/src"
	"fmt"
)

func main() {
	arg1 := struct{ Name string }{}
	fmt.Println("has name", src.HasField(arg1, "name"))
	fmt.Println("has age", src.HasField(arg1, "age"))

	arg2 := &struct {
		Random int
		Owner  *struct {
			Address struct {
				Name string
				Cp   int
			}
		}
	}{}

	fmt.Println("has random", src.HasField(arg2, "random"))
	fmt.Println("has owner.address", src.HasField(arg2, "owner.address"))
	fmt.Println("has owner.address.age", src.HasField(arg2, "owner.address.age"))
	fmt.Println("has owner.address.cp", src.HasField(arg2, "owner.address.cp"))
}

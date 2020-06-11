package main

import (
	"carlware/hasfield/src"
	"fmt"
)

func main() {
	arg1 := struct{ Name string }{}
	fmt.Println("has name", src.HasField(arg1, "name"))
	fmt.Println("has age", src.HasField(arg1, "age"))

}
